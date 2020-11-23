/*
 * PoGoStack Compiler www.pogostack.com
 *
 * Author: "Alex Nedoboi" alex@pogostack.com
 *
 * Contributors:
 *
 * "Evgeny Komarevtsev" evgeny@komarevtsev.net
 *
 * Input: .pogo file format, see pogostack.org for specifications and examples
 *
 * Output: PL/pgSQL function source, ready to pipe to psql
 *
 * Usage: go run pogo_compiler.go input_file_name.pogo | psql -h host -d database -U user
 *
 */

package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var compilerVersion = "pogo1.7.1"
var BUILD_VERSION = ""
var functionName, friendlyFunctionName, friendlyFunctionNameOnly, functionPrefix, functionParameters, functionReturns, functionForm, templateSuffix, sqlDebug, volatilityCategory string
var testcases, dependencies, nativeParameters []string
var isNoauth bool
var isPCheck bool = true
var isDebuggable bool
var isTraceEnabled bool
var isTraceParametersInDOM bool
var breakpointCount = 0
var debuggerBreakpoints []int
var pogoFileName = ""
var rootCodePath = ""
var relativeFolderHash = ""
var relativeFolderName = ""

func writePogoBreakPoint(b *bytes.Buffer, debugVariables []string, lineNumber int) {

	if !isDebuggable {
		return
	}

	b.WriteString(fmt.Sprintf("\n"))
	b.WriteString(fmt.Sprintf("loop \n"))
	{
		for _, variable := range debugVariables {
			b.WriteString(fmt.Sprintf("_ds_ := _ds_ || jsonb_build_object('%v', %v);", variable, variable))
		}

		if breakpointCount > 0 {
			b.WriteString(fmt.Sprintf("perform __pogo_stack_push('%v', %v, %v, _ds_, _thread_id_, '%v', _frame_depth_);", friendlyFunctionNameOnly, lineNumber, true, pogoFileName))
		} else {
			b.WriteString(fmt.Sprintf("_frame_depth_ := __pogo_stack_push('%v', %v, %v, _ds_, _thread_id_, '%v', null::integer);", friendlyFunctionNameOnly, lineNumber, false, pogoFileName))
		}

		b.WriteString(fmt.Sprintf("if __pogo_break_point_should_stop(%v, '%v', _thread_id_, _frame_depth_, %v) then \n", lineNumber, friendlyFunctionNameOnly, isTraceEnabled))
		{
			b.WriteString(fmt.Sprintf("_ds_bs_ := __pogo_break_point(%v, '%v', _thread_id_, _frame_depth_);\n", lineNumber, friendlyFunctionNameOnly))
			b.WriteString(fmt.Sprintf(`case _ds_bs_->>'command'
                                            when 'retry' then continue;
                                            when 'continue' then exit;
                                            else exit;
                                        end case;
                                        `))
			b.WriteString(fmt.Sprintf(`else exit;`))
		}
		b.WriteString(fmt.Sprintf("end if;\n"))
	}
	b.WriteString(fmt.Sprintf("end loop; \n"))

	debuggerBreakpoints = append(debuggerBreakpoints, lineNumber)
	breakpointCount++

}

func importFile(fileName string) string {

	embeddedFile, err := Asset(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open %v\n\n", fileName)
		os.Exit(1)
	}

	var b bytes.Buffer
	line := ""
	file := bytes.NewReader(embeddedFile)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.Replace(scanner.Text(), `$name$`, functionName, -1)
		line = strings.Replace(line, `$friendly_name$`, friendlyFunctionName, -1)
		line = strings.Replace(line, `$function$`, functionPrefix+functionName, -1)

		isNoauthSQLBool := "false"
		if isNoauth {
			isNoauthSQLBool = "true"
		}

		line = strings.Replace(line, `$is_noauth$`, isNoauthSQLBool, -1)

		parametersLen := len(functionParameters)
		if parametersLen > 0 && functionParameters[parametersLen-1] == ',' {
			functionParameters = functionParameters[:parametersLen-1]
		}
		line = strings.Replace(line, `$parameters$`, functionParameters, -1)
		line = strings.Replace(line, `$returns$`, functionReturns, -1)
		line = strings.Replace(line, `$form$`, functionForm, -1)
		line = strings.Replace(line, `$volatility_category$`, volatilityCategory, -1)

		st := "null"
		if testcases != nil {
			st = `'` + strings.Join(testcases, ",") + `'`
		}
		line = strings.Replace(line, "$testcases$", st, -1)

		st = "null"
		if dependencies != nil {
			st = `'` + strings.Join(dependencies, ",") + `'`
		}
		line = strings.Replace(line, `$dependencies$`, st, -1)

		line = strings.Replace(line, "$debug$", sqlDebug, -1)

		if isDebuggable && breakpointCount > 0 {
			var values []string

			for _, breakpoint := range debuggerBreakpoints {
				values = append(values, fmt.Sprintf("('%v', %v)", friendlyFunctionNameOnly, breakpoint))
			}

			line = strings.Replace(line, "$debugger_breakpoints$", fmt.Sprintf("insert into __pogo__breakpoints (page, line) values %v;", strings.Join(values, ",")), -1)
		} else {
			line = strings.Replace(line, "$debugger_breakpoints$", "", -1)
		}

		b.WriteString(line + "\n")
	}

	var hh = b.String()

	return hh
}

func main() {

	build := flag.String("b", "", "Makes $BUILD$ available in pogo files")
	isPrintVersion := flag.Bool("v", false, "Displays compiler version and exits")
	isCleanup := flag.Bool("cleanup", false, "Clean up (remove previously compiled but do not compile new version)")
	isDebuggableFlag := flag.Bool("debug", false, "Enables debug build (breakpoint support)")
	isTraceFlag := flag.Bool("trace", false, "Enables tracing into __pogo_debugger_trace (works with debugging enabled only)")
	pTraceFlag := flag.Bool("ptrace", false, "Enables output of function parameters in DOM rendering)")
	srcRootPathFlag := flag.String("root", "", "Root directory code file is compiled against")
	flag.Parse()

	isDebuggable = *isDebuggableFlag
	isTraceEnabled = *isTraceFlag
	isTraceParametersInDOM = *pTraceFlag
	rootCodePath = *srcRootPathFlag

	if *isPrintVersion {
		fmt.Fprintf(os.Stdout, "pogo_compiler version %v\n", compilerVersion)
		os.Exit(0)
	}

	BUILD_VERSION = *build

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: pogo_compiler -b [build_version] input_file.pogo")
		os.Exit(1)
	}

	pogoFileName, _ = filepath.Abs(os.Args[len(os.Args)-1])

	p2, _ := filepath.Split(pogoFileName)
	relPath, _ := filepath.Rel(filepath.Clean(strings.ToLower(rootCodePath)), filepath.Clean(strings.ToLower(p2)))
	relPath = strings.Replace(relPath, "\\", "/", -1)
	relativeFolderHash = fmt.Sprintf("%x_", md5.Sum([]byte(relPath)))
	relativeFolderName = relPath + "/"

	if relPath == "." || rootCodePath == "" {
		relativeFolderHash = ""
		relativeFolderName = ""
	}

	b := compileFile(pogoFileName, isCleanup)

	fmt.Println(b.String())

	os.Exit(0)

}

func compileFile(pogoFileName string, isCleanup *bool) bytes.Buffer {
	sourceFile, err := os.Open(pogoFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open input file %v \n\n", pogoFileName)
		os.Exit(1)
	}
	defer sourceFile.Close()
	inCode := false
	inEquals := false
	inDeclare := false
	inHeader := true
	inParams := false
	inSQLComment := false
	inHTMLComment := false
	inJSONB := false
	inCoalesce := false
	inForm := false
	jsonb := 0
	form := 0
	inQuotesSingle := false
	inQuotesDouble := false
	inQuotesBacktick := false
	toSkip := false
	tagProcessed := false
	isFirstLine := true
	isIncluding := false
	isNative := false
	isRaw := false
	isJSONB := false
	isView := false
	isPragmaTiming := false
	inNewDeclareStatement := false
	lineInput := ""
	var b bytes.Buffer
	tags := []string{}
	tagsToIgnore := map[string]bool{
		"/area":   true,
		"/base":   true,
		"/br":     true,
		"/col":    true,
		"/embed":  true,
		"/hr":     true,
		"/img":    true,
		"/input":  true,
		"/link":   true,
		"/meta":   true,
		"/param":  true,
		"/source": true,
		"/track":  true,
		"/wbr":    true,
		"/path":   true,
	}
	sqlDebug = "1"
	functionPrefix = "psp2_"
	functionParameters = ""
	functionReturns = "pogo__return__type"
	volatilityCategory = "VOLATILE"
	isNative = true
	nativeParameters = []string{}
	functionForm = "'{}'"
	reSpaces := regexp.MustCompile("[ ]{2,}")
	rePSPFull := regexp.MustCompile(" psp_[a-z_]{1,}")
	rePSPName := regexp.MustCompile("^ psp_")
	reTag := regexp.MustCompile("^[/]{0,1}[a-z]{1,}[0-9]{0,1}")
	debugVariables := []string{}
	//is_debuggable := true
	scanner := bufio.NewScanner(sourceFile)
	lineNumber := 0
	for (!isIncluding && scanner.Scan()) || isIncluding {

		lineNumber++
		lineInput = scanner.Text()

		if strings.HasPrefix(lineInput, " ") || regexp.MustCompile("^[\t]+[ ]").MatchString(lineInput) {
			fmt.Fprintf(os.Stderr, "\x1b[31;1mBad indent in `%v` line %v \x1b[0m \n", functionName, lineNumber)
			os.Exit(1)
		}

		if !inQuotesBacktick {
			lineInput = strings.Replace(lineInput, "\t", " ", -1)
			lineInput = strings.Trim(lineInput, " ")
			lineInput = reSpaces.ReplaceAllString(lineInput, " ")
		}

		lineInput = strings.Replace(lineInput, "$function_name$", functionName, -1)

		if lineInput == "" && !inQuotesBacktick {
			continue
		}

		if strings.HasPrefix(lineInput, "//") {
			continue
		}

		/* comments section in the beginning of .pogo file, only relevant until the first non-comment line */

		if isFirstLine {

			if lineInput == "#pragma noauth" {
				isNoauth = true
				continue
			}

			if lineInput == "#pragma nopcheck" {
				isPCheck = false
				continue
			}

			if strings.HasPrefix(lineInput, "#testcase ") {
				splices := strings.SplitN(lineInput+" ", " ", 2)
				testcase := strings.Trim(splices[1], " ")
				if testcase != "" {
					testcases = append(testcases, testcase)
				}
				continue
			}

			if lineInput == "#pragma nowrap" {
				sqlDebug = "0"
				continue
			}

			if lineInput == "#pragma raw" {
				isRaw = true
				functionPrefix = "f_"
				functionParameters = ""
				//native_parameters = []
				templateSuffix = "_raw"
				continue
			}

			if lineInput == "#pragma immutable" {
				volatilityCategory = "IMMUTABLE"
				isDebuggable = false
				continue
			}

			if lineInput == "#pragma stable" {
				volatilityCategory = "STABLE"
				continue
			}

			if lineInput == "#pragma timing" {
				isPragmaTiming = true
				continue
			}

			if lineInput == "#pragma nodebug" {
				isDebuggable = false
				continue
			}

			if strings.HasPrefix(lineInput, "#returns ") {
				splices := strings.SplitN(lineInput, " ", 2)
				functionReturns = splices[1]
				continue
			}

			if strings.HasPrefix(lineInput, "#") {
				continue
			}

			if relativeFolderName == "" || strings.HasPrefix(relativeFolderName, "pogo_system/") {
				functionName = lineInput
			} else {
				functionName = fmt.Sprintf("%x", md5.Sum([]byte(relativeFolderName+lineInput)))
			}
			friendlyFunctionName = relativeFolderName + lineInput
			friendlyFunctionNameOnly = lineInput

			if !isView {
				b.WriteString(importFile("templates/template_psp_function_drop.sql"))
				if isDebuggable {
					b.WriteString(fmt.Sprintf("delete from __pogo__breakpoints where page = '%v';\n", friendlyFunctionNameOnly))
				}
			}

			if *isCleanup {
				os.Exit(0)
			}

			if !((isNative || isRaw || isView) && !isJSONB) {
				b.WriteString(importFile("templates/template_psp_function_begin" + templateSuffix + ".sql"))
			}

			isFirstLine = false
			inParams = true

		} else {

			if psps := rePSPFull.FindAllString(lineInput, -1); psps != nil {
				for _, psp := range psps {
					dependencies = append(dependencies, rePSPName.ReplaceAllString(psp, ""))
				}
			}

			if rootCodePath == "" {
				lineInput = strings.Replace(lineInput, "<%= psp2_", "<= (select text_content from psp2_", -1)
			} else {
				preReplace := lineInput
				lineInput = strings.Replace(lineInput, "<%= psp2/", "<= (select text_content from psp2_", -1)
				if preReplace != lineInput { //need to call a psp2 including path
					s := strings.Index(preReplace, "psp2/")
					if s == -1 {
						continue
					}
					newS := preReplace[s+len("psp2/"):]
					e := strings.Index(newS, "(")
					if e == -1 {
						continue
					}
					urlPath := newS[:e]
					tokenToReplace := "<%= psp2/" + urlPath
					urlPath = strings.TrimSpace(filepath.Clean(urlPath))
					pathMD5 := ""
					if urlPath == "." {
						urlPath = ""
					} else {
						urlPath = strings.Replace(urlPath, "\\", "/", -1)
						pathMD5 = fmt.Sprintf("%x", md5.Sum([]byte(urlPath)))
					}

					replacementString := fmt.Sprintf("<= (select text_content from psp2_%v", pathMD5)
					lineInput = strings.Replace(preReplace, tokenToReplace, replacementString, -1)
				}
			}
			lineInput = strings.Replace(lineInput, "pogo_json(", "_jsonr_ := (", -1)
			lineInput = strings.Replace(lineInput, "pogo_binary(", "_bytear_ := (", -1)
			lineInput = strings.Replace(lineInput, "pogo_http_code(", "_addr_ := _addr_ || jsonb_build_object('http_code', ", -1)
			lineInput = strings.Replace(lineInput, "pogo_header(", "_addr_ := jsonb_set(_addr_, '{headers}',", -1)
			lineInput = strings.Replace(lineInput, "pogo_cookie(", "_addr_ := jsonb_set(_addr_, '{cookies}',", -1)

			/* variables declaration tags <? ?> */
			if strings.HasPrefix(lineInput, "<?") {
				inHeader = false /* as soon as we reach the declare section, the header section stops */
				inParams = false
				inDeclare = true
				inNewDeclareStatement = true /* that is used to parse variable names in declaration */
				lineInput = lineInput[2:]

				if (isNative || isRaw || isView) && !isJSONB {
					b.WriteString(importFile("templates/template_psp_function_begin" + templateSuffix + ".sql"))
				}
			}

			if strings.HasPrefix(lineInput, "?>") {
				inDeclare = false

				if !isView {
					b.WriteString("begin\n")
				}

				if !isRaw && !isView {

					traceParameters := ""

					if isTraceParametersInDOM {
						traceParameters = `' || replace(__FORM__, '&', ' ') || '`
						if isNative || isRaw {
							traceParameters = ``
							if isJSONB {
								traceParameters = `' || jsonb_pretty(p_jsonb) || '`
							} else {
								for _, nativeParameter := range nativeParameters {
									traceParameters += nativeParameter + ` := ' || coalesce(` + nativeParameter + `::varchar, '<null>') || ' | `
								}
							}
						}
					}

					b.WriteString("if _d_ > 0 then _v_[_n_] := '<!-- start: " + friendlyFunctionName + " " + traceParameters + " -->'; _n_ := _n_ + 1; end if;\n_v_[_n_] := '")
				}
				lineInput = lineInput[2:]
			}

			/* parameters */
			if inParams {

				/*
					Input  : parameter type default
					Output : parameter type := pgasp_parse_get[_type] (__FORM__|__CONTEXT__, 'parameter', 'default');
				*/

				if strings.Count(lineInput, " ") < 2 {
					lineInput += " "
				}

				splice := strings.SplitN(lineInput, " ", 3)

				if isPCheck && !strings.HasPrefix(splice[0], "p_") && !strings.HasPrefix(splice[0], "sp_") && !strings.HasPrefix(splice[0], "new_") && len(splice[0]) != 1 {
					fmt.Fprintf(os.Stderr, "\x1b[31;1mBad parameter name `%v` in `%v` \x1b[0m \n", splice[0], functionName)
					os.Exit(1)
				}

				if splice[0] != "" {
					debugVariables = append(debugVariables, splice[0])
				}

				if (isNative || isRaw) && !isJSONB {

					functionParameters += fmt.Sprintf("%s %s default %s,", splice[0], splice[1], splice[2])
					nativeParameters = append(nativeParameters, splice[0])

				} else {

					if !isJSONB {

						b.WriteString(fmt.Sprintf("%s %s := pgasp_parse_get", splice[0], splice[1]))

						if splice[1] == "date" || splice[1] == "json" {
							b.WriteString("_" + splice[1])
						}

						source := "__FORM__"
						if splice[2] == "#context" {
							source = "__CONTEXT__"
						}

						b.WriteString(fmt.Sprintf("(%s, '%s', '%s');\n", source, splice[0], splice[2]))

					} else {

						b.WriteString(fmt.Sprintf("%s %s := coalesce((p_jsonb ->> '%s')::%s, %s);\n", splice[0], splice[1], splice[0], splice[1], splice[2]))

					}

				}

				continue
			}

			if inDeclare {
				constantParameterName := regexp.MustCompile(`^[A-Z_]{1,}`).FindString(lineInput)

				if constantParameterName != "" {
					debugVariables = append(debugVariables, constantParameterName)
					inNewDeclareStatement = false
				} else {
					declareSplice := strings.SplitN(strings.TrimRight(lineInput, ";"), " ", 3)
					splice := declareSplice[0]
					if inNewDeclareStatement && splice != "" && len(declareSplice) >= 2 && declareSplice[1] != "record" {
						debugVariables = append(debugVariables, splice)
						inNewDeclareStatement = false
					}
				}
				if strings.HasSuffix(strings.TrimSpace(lineInput), ";") {
					inNewDeclareStatement = true
				}

				if constantParameterName != "" && strings.Index(lineInput, `:=`) == -1 {
					constantParameterType := strings.Trim(regexp.MustCompile(` [A-Za-z]{1,}`).FindString(lineInput), " ")
					b.WriteString(fmt.Sprintf("%s constant %s := get_system_value_%s('%s');\n", constantParameterName, constantParameterType, constantParameterType, strings.ToLower(constantParameterName)))
					continue
				}
			}

			/* function name, passed parameters, and declared variables - done, now processing the rest */

			lineBytes := append([]byte(lineInput), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}...)

			i := 0

			for lineBytes[i] > 0 {

				if string(lineBytes[i:i+7]) == `$BUILD$` {
					var buildVersion = BUILD_VERSION
					if isDebuggable {
						buildVersion += "*"
					}
					b.WriteString(buildVersion)
					i += 7
				}

				tagProcessed = true

				for tagProcessed {

					tagProcessed = false /* to cover 2+ consecutive tags */

					/* comments */

					if (inCode || inEquals) && lineBytes[i] == '/' && lineBytes[i+1] == '*' {
						inSQLComment = true
						i += 2
					}

					if (inCode || inEquals) && inSQLComment && lineBytes[i] == '*' && lineBytes[i+1] == '/' {
						inSQLComment = false
						i += 2
					}

					if !inCode && !inEquals {

						if string(lineBytes[i:i+4]) == `<!--` {
							inHTMLComment = true
							i += 4
						}

						if string(lineBytes[i:i+3]) == `-->` && inHTMLComment {
							inHTMLComment = false
							i += 3
						}

						if string(lineBytes[i:i+5]) == "<form" {
							inForm = true
						}

						if lineBytes[i] == '<' {
							tag := reTag.FindString(string(lineBytes[i+1:]))
							if tag != "" && !(tagsToIgnore[tag] || tagsToIgnore["/"+tag]) {
								tagsLength := len(tags)
								if strings.HasPrefix(tag, "/") {
									if tagsLength == 0 {
										fmt.Fprintf(os.Stderr, "\x1b[31;1mCannot close un-opened tag <%v> \x1b[0m \nFile: %v \nLine: %v \nTags: %v \n\n",
											tag, friendlyFunctionName, lineNumber, tags)
										os.Exit(1)
									}
									if "/"+tags[tagsLength-1] != tag {
										fmt.Fprintf(os.Stderr, "\x1b[31;1mClosing tag <%v> mismatches opened tag <%v> \x1b[0m \nFile: %v \nLine: %v \nTags: %v \n\n",
											tag, tags[tagsLength-1], friendlyFunctionName, lineNumber, tags)
										os.Exit(1)
									}
									tags = tags[:tagsLength-1]
								} else {
									tags = append(tags, tag)
								}
								//fmt.Fprintf(os.Stderr, "tags at line %v: %v\n", line_number, tags)
							}
						}

					}

					/* quotes */

					if inCode || inEquals || inDeclare {

						if !inQuotesSingle && !inQuotesBacktick && lineBytes[i] == '"' {
							inQuotesDouble = !inQuotesDouble
						}

						if !inQuotesDouble && !inQuotesBacktick && lineBytes[i] == '\'' {
							inQuotesSingle = !inQuotesSingle
						}

						if !inQuotesSingle && !inQuotesDouble && lineBytes[i] == '`' {
							inQuotesBacktick = !inQuotesBacktick
						}
					}

					if (inCode || inEquals || inDeclare) && !(inQuotesSingle || inQuotesDouble || inQuotesBacktick) {

						if lineBytes[i] == '{' {
							inJSONB = true
						}

						if lineBytes[i] == '!' && lineBytes[i+1] != '=' && lineBytes[i+1] != '~' {
							inCoalesce = true
							b.WriteString(`devnull(`)
							i++
							tagProcessed = true
						}

						if lb := lineBytes[i]; inCoalesce {
							if !((lb >= 'a' && lb <= 'z') || (lb >= '0' && lb <= '9') || lb == '_' || lb == '.') {
								inCoalesce = false
								b.WriteString(`)`)
							}
						}

						if lineBytes[i] == '.' && lineBytes[i+1] == '.' && lineBytes[i+2] == '.' && lineBytes[i+3] == '{' {
							b.WriteString(`p_jsonb||`)
							i += 3
							tagProcessed = true
						}

					}

					/* jsonb */

					if inJSONB /*&& !(in_quotes_single || in_quotes_double || in_quotes_backtick)*/ {

						isBuildingJSONB := true

						if lineBytes[i] == '\\' && lineBytes[i+1] == '{' {
							b.WriteString(`{`)
							i++
						} else if lineBytes[i] == '\\' && lineBytes[i+1] == '}' {
							b.WriteString(`}`)
							i++
						} else {

							switch lineBytes[i] {
							case '{':
								b.WriteString(`jsonb_build_object(`)
								jsonb++
							case '}':
								b.WriteString(`)`)
								jsonb--
								if jsonb == 0 {
									inJSONB = false
								}
							case '[':
								b.WriteString(`jsonb_build_array(`)
							case ']':
								b.WriteString(`)`)
							case ':':
								if lineBytes[i-1] != ':' && lineBytes[i+1] != ':' && lineBytes[i+1] != '=' {
									b.WriteString(`,`)
								} else {
									isBuildingJSONB = false
								}
							case '"':
								b.WriteString(`'`)
							default:
								isBuildingJSONB = false
							}
						}
						if isBuildingJSONB {
							i++
							tagProcessed = true

						}
					}

					/* code tag <% %> but not print tag <%= %> */

					if !inCode && lineBytes[i] == '<' && lineBytes[i+1] == '%' && lineBytes[i+2] != '=' {
						inCode = true
						if !isRaw && !isView {
							b.WriteString(fmt.Sprintf("'; _n_ := _n_ + 1; _l_ := %v; ", lineNumber))

							writePogoBreakPoint(&b, debugVariables, lineNumber)

							if isPragmaTiming {
								b.WriteString("_tm_ := _tm_ || jsonb_build_object('Line ' || _l_ || ':', ((clock_timestamp()-_tc_)::interval)); _tc_ := clock_timestamp(); ")
							}
						}

						i += 2
						tagProcessed = true
					}

					if lineBytes[i] == '%' && lineBytes[i+1] == '>' {

						if inCode {
							inCode = false
							if !isRaw && !isView {
								b.WriteString(`_v_[_n_] := '`)
							}
							i += 2
							tagProcessed = true
						}

						if inEquals {
							inEquals = false
							b.WriteString(`)) || '`)
							i += 2
							tagProcessed = true
						}
					}

					/* new style print tag <= => */
					if !inCode && !inEquals && lineBytes[i] == '<' && lineBytes[i+1] == '=' {
						inEquals = true
						b.WriteString(`' || (`)
						i += 2
						tagProcessed = true
					}

					/* classic style print tag <%= %> */
					if !inCode && lineBytes[i] == '<' && lineBytes[i+1] == '%' && lineBytes[i+2] == '=' {
						inEquals = true
						b.WriteString(`' || coalesce2('' || (`)
						i += 3
						tagProcessed = true
					}

					if !inCode && !inEquals {

						if inForm && lineBytes[i] == '<' {
							form++
						}

						if inForm && lineBytes[i] == '>' {
							form--
							// if form == 0 {
							// 	b.WriteByte(line_bytes[i])
							// 	line_bytes = append([]byte(`<input type="hidden" name="new_session_token" value="<%= f_get_session_token(CU) %>">`), line_bytes[i+1:]...)
							// 	i = 0
							// 	in_form = false
							// }
						}
					}

					if (inCode || inEquals) && lineBytes[i] == '`' {
						b.WriteString(`$tick$`)
						i++
					}

				} /* for tag_processed */

				/* removing comments, both pgsql and html */
				toSkip = false

				if (inCode || inEquals) && inSQLComment {
					toSkip = true
				}

				if (!inCode && !inEquals) && inHTMLComment {
					toSkip = true
				}

				if !toSkip {

					if !inCode && !inEquals && !inDeclare && !inHeader && lineBytes[i] == '\'' {
						b.WriteByte(lineBytes[i])
					}

					if lineBytes[i] > 0 {
						b.WriteByte(lineBytes[i])
					}

					if inCode && !isView && !inQuotesBacktick && !inQuotesDouble && !inQuotesSingle && !inJSONB && lineBytes[i] == ';' {
						writePogoBreakPoint(&b, debugVariables, lineNumber)
					}
				}

				i++

			} /* regular line */

		} /* not first line */

		b.WriteString("\n")

	}
	/* for scanner.scan */
	if len(tags) > 0 {
		fmt.Fprintf(os.Stderr, "\x1b[31;1mOpened tags not closed \x1b[0m \nFile: %v \nTags: %v \n\n",
			friendlyFunctionName, tags)
		os.Exit(1)
	}
	if !isRaw && !isView {
		b.WriteString("';\n_n_ := _n_ + 1;\n")
	}
	if rootCodePath == "" {
		b.WriteString(importFile("templates/template_psp_function_end" + templateSuffix + ".sql"))
	} else {
		b.WriteString(importFile("templates/template_psp_function_end_routed" + templateSuffix + ".sql"))
	}
	return b
} /* main */
