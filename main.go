/*
 * PoGoStack Compiler www.pogostack.com
 *
 * Author: "Alex Nedoboi" alex@pogostack.com
 * Contributo: "Evgeny Komarevtsev" evgeny@komarevtsev.net
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
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var compiler_version = "pogo1.0"
var function_name, function_prefix, function_parameters, function_returns, function_form, template_suffix, sql_debug, volatility_category string
var testcases, dependencies, native_parameters []string
var is_noauth bool
var BUILD_VERSION = ""
var is_debuggable bool
var is_trace_enabled bool
var breakpoint_count = 0
var debugger_breakpoints []int
var pogo_file_name = ""

func writePogoBreakPoint(b *bytes.Buffer, debug_variables []string, lineNumber int) {

	if !is_debuggable {
		return
	}

	b.WriteString(fmt.Sprintf("\n"))
	b.WriteString(fmt.Sprintf("loop \n"))
	{
		for _, variable := range debug_variables {
			b.WriteString(fmt.Sprintf("_ds_ := _ds_ || jsonb_build_object('%v', %v);", variable, variable))
		}

		if breakpoint_count > 0 {
			b.WriteString(fmt.Sprintf("perform __pogo_stack_push('%v', %v, %v, _ds_, _thread_id_, '%v', _frame_depth_);", function_name, lineNumber, true, pogo_file_name))
		} else {
			b.WriteString(fmt.Sprintf("_frame_depth_ := __pogo_stack_push('%v', %v, %v, _ds_, _thread_id_, '%v', null::integer);", function_name, lineNumber, false, pogo_file_name))
		}

		b.WriteString(fmt.Sprintf("if __pogo_break_point_should_stop(%v, '%v', _thread_id_, _frame_depth_, %v) then \n", lineNumber, function_name, is_trace_enabled))
		{
			b.WriteString(fmt.Sprintf("_ds_bs_ := __pogo_break_point(%v, '%v', _thread_id_, _frame_depth_);\n", lineNumber, function_name))
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

	debugger_breakpoints = append(debugger_breakpoints, lineNumber)
	breakpoint_count++

}

func importFile(file_name string) string {

	embeddedFile, err := Asset(file_name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open %v\n\n", file_name)
		os.Exit(1)
	}

	var b bytes.Buffer
	line := ""
	file := bytes.NewReader(embeddedFile)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.Replace(scanner.Text(), `$name$`, function_name, -1)
		line = strings.Replace(line, `$function$`, function_prefix+function_name, -1)

		is_noauth_sql_bool := "false"
		if is_noauth {
			is_noauth_sql_bool = "true"
		}

		line = strings.Replace(line, `$is_noauth$`, is_noauth_sql_bool, -1)

		parameters_len := len(function_parameters)
		if parameters_len > 0 && function_parameters[parameters_len-1] == ',' {
			function_parameters = function_parameters[:parameters_len-1]
		}
		line = strings.Replace(line, `$parameters$`, function_parameters, -1)
		line = strings.Replace(line, `$returns$`, function_returns, -1)
		line = strings.Replace(line, `$form$`, function_form, -1)
		line = strings.Replace(line, `$volatility_category$`, volatility_category, -1)

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

		line = strings.Replace(line, "$debug$", sql_debug, -1)

		if is_debuggable && breakpoint_count > 0 {
			var values []string

			for _, breakpoint := range debugger_breakpoints {
				values = append(values, fmt.Sprintf("('%v', %v)", function_name, breakpoint))
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
	print_version := flag.Bool("v", false, "Displays compiler version and exits")
	environment := flag.String("e", "p", "Environment [d|t|p]")
	is_cleanup := flag.Bool("cleanup", false, "Clean up (remove previously compiled but do not compile new version)")
	is_debuggable_flag := flag.Bool("debug", false, "Enables debug build (breakpoint support)")
	is_trace_flag := flag.Bool("trace", false, "Enables tracing into __pogo_debugger_trace (works with debugging enabled only)")
	flag.Parse()

	is_debuggable = *is_debuggable_flag
	is_trace_enabled = *is_trace_flag

	if *print_version {
		fmt.Fprintf(os.Stdout, "pogo_compiler version %v\n", compiler_version)
		os.Exit(0)
	}

	BUILD_VERSION = *build

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: pogo_compiler -b [build_version] input_file.pogo")
		os.Exit(1)
	}

	pogo_file_name, _ = filepath.Abs(os.Args[len(os.Args)-1])
	source_file, err := os.Open(pogo_file_name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open input file %v \n\n", pogo_file_name)
		os.Exit(1)
	}

	defer source_file.Close()

	in_code := false
	in_equals := false
	in_declare := false
	in_header := true
	in_params := false
	in_sql_comment := false
	in_html_comment := false
	in_jsonb := false
	in_coalesce := false
	in_form := false
	jsonb := 0
	form := 0
	in_quotes_single := false
	in_quotes_double := false
	in_quotes_backtick := false
	to_skip := false
	tag_processed := false
	is_first_line := true
	is_including := false
	is_native := false
	is_raw := false
	is_jsonb := false
	is_view := false
	is_pragma_timing := false
	in_new_declare_statement := false

	line_input := ""
	var b bytes.Buffer

	tags := []string{}
	tags_to_ignore := map[string]bool{
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
	}

	sql_debug = "1"

	function_prefix = "psp2_"
	function_parameters = ""
	function_returns = "pogo__return__type"
	volatility_category = "VOLATILE"

	is_native = true
	native_parameters = []string{}
	function_form = "'{}'"

	re_spaces := regexp.MustCompile("[ ]{2,}")
	re_psp_full := regexp.MustCompile(" psp_[a-z_]{1,}")
	re_psp_name := regexp.MustCompile("^ psp_")
	re_tag := regexp.MustCompile("^[/]{0,1}[a-z]{1,}[0-9]{0,1}")

	debug_variables := []string{}
	//is_debuggable := true

	scanner := bufio.NewScanner(source_file)

	line_number := 0

	for (!is_including && scanner.Scan()) || is_including {

		line_number++
		line_input = scanner.Text()

		if strings.HasPrefix(line_input, " ") || regexp.MustCompile("^[\t]+[ ]").MatchString(line_input) {
			fmt.Fprintf(os.Stderr, "\x1b[31;1mBad indent in `%v` line %v \x1b[0m \n", function_name, line_number)
			os.Exit(1)
		}

		line_input = strings.Replace(line_input, "\t", " ", -1)
		line_input = strings.Trim(line_input, " ")
		line_input = re_spaces.ReplaceAllString(line_input, " ")

		line_input = strings.Replace(line_input, "$function_name$", function_name, -1)

		if line_input == "" {
			continue
		}

		if strings.HasPrefix(line_input, "//") {
			continue
		}

		/* comments section in the beginning of .pogo file, only relevant until the first non-comment line */

		if is_first_line {

			if line_input == "#pragma noauth" {
				is_noauth = true
				continue
			}

			if strings.HasPrefix(line_input, "#testcase ") {
				splices := strings.SplitN(line_input+" ", " ", 2)
				testcase := strings.Trim(splices[1], " ")
				if testcase != "" {
					testcases = append(testcases, testcase)
				}
				continue
			}

			if line_input == "#pragma nowrap" {
				sql_debug = "0"
				continue
			}

			if line_input == "#pragma raw" {
				is_raw = true
				function_prefix = "f_"
				function_parameters = ""
				//native_parameters = []
				template_suffix = "_raw"
				continue
			}

			if line_input == "#pragma immutable" {
				volatility_category = "IMMUTABLE"
				is_debuggable = false
				continue
			}

			if line_input == "#pragma stable" {
				volatility_category = "STABLE"
				continue
			}

			if line_input == "#pragma timing" {
				is_pragma_timing = true
				continue
			}

			if line_input == "#pragma nodebug" {
				is_debuggable = false
				continue
			}

			if strings.HasPrefix(line_input, "#returns ") {
				splices := strings.SplitN(line_input, " ", 2)
				function_returns = splices[1]
				continue
			}

			if strings.HasPrefix(line_input, "#") {
				continue
			}

			function_name = line_input

			if !is_view {
				b.WriteString(importFile("templates/template_psp_function_drop.sql"))
				if is_debuggable {
					b.WriteString(fmt.Sprintf("delete from __pogo__breakpoints where page = '%v';\n", function_name))
				}
			}

			if *is_cleanup {
				os.Exit(0)
			}

			if !((is_native || is_raw || is_view) && !is_jsonb) {
				b.WriteString(importFile("templates/template_psp_function_begin" + template_suffix + ".sql"))
			}

			is_first_line = false
			in_params = true

		} else {

			if psps := re_psp_full.FindAllString(line_input, -1); psps != nil {
				for _, psp := range psps {
					dependencies = append(dependencies, re_psp_name.ReplaceAllString(psp, ""))
				}
			}

			line_input = strings.Replace(line_input, "<%= psp2_", "<= (select text_content from psp2_", -1)
			line_input = strings.Replace(line_input, "pogo_json(", "_jsonr_ := (", -1)
			line_input = strings.Replace(line_input, "pogo_binary(", "_bytear_ := (", -1)
			line_input = strings.Replace(line_input, "pogo_http_code(", "_addr_ := _addr_ || jsonb_build_object('http_code', ", -1)
			line_input = strings.Replace(line_input, "pogo_header(", "_addr_ := jsonb_set(_addr_, '{headers}',", -1)
			line_input = strings.Replace(line_input, "pogo_cookie(", "_addr_ := jsonb_set(_addr_, '{cookies}',", -1)

			/* variables declaration tags <? ?> */
			if strings.HasPrefix(line_input, "<?") {
				in_header = false /* as soon as we reach the declare section, the header section stops */
				in_params = false
				in_declare = true
				in_new_declare_statement = true /* that is used to parse variable names in declaration */
				line_input = line_input[2:]

				if (is_native || is_raw || is_view) && !is_jsonb {
					b.WriteString(importFile("templates/template_psp_function_begin" + template_suffix + ".sql"))
				}
			}

			if strings.HasPrefix(line_input, "?>") {
				in_declare = false

				if !is_view {
					b.WriteString("begin\n")
				}

				if !is_raw && !is_view {

					traceParameters := ""

					if *environment == "d" {
						traceParameters = `' || replace(__FORM__, '&', ' ') || '`
						if is_native || is_raw {
							traceParameters = ``
							if is_jsonb {
								traceParameters = `' || jsonb_pretty(p_jsonb) || '`
							} else {
								for _, native_parameter := range native_parameters {
									traceParameters += native_parameter + ` := ' || coalesce(` + native_parameter + `::varchar, '<null>') || ' | `
								}
							}
						}
					}

					b.WriteString("if _d_ > 0 then _v_[_n_] := '<!-- start: " + function_name + " " + traceParameters + " -->'; _n_ := _n_ + 1; end if;\n_v_[_n_] := '")
				}
				line_input = line_input[2:]
			}

			/* parameters */
			if in_params {

				/*
					Input  : parameter type default
					Output : parameter type := pgasp_parse_get[_type] (__FORM__|__CONTEXT__, 'parameter', 'default');
				*/

				if strings.Count(line_input, " ") < 2 {
					line_input += " "
				}

				splice := strings.SplitN(line_input, " ", 3)

				if !strings.HasPrefix(splice[0], "p_") && !strings.HasPrefix(splice[0], "sp_") && !strings.HasPrefix(splice[0], "new_") && len(splice[0]) != 1 {
					fmt.Fprintf(os.Stderr, "\x1b[31;1mBad parameter name `%v` in `%v` \x1b[0m \n", splice[0], function_name)
					os.Exit(1)
				}

				if splice[0] != "" {
					debug_variables = append(debug_variables, splice[0])
				}

				if (is_native || is_raw) && !is_jsonb {

					function_parameters += fmt.Sprintf("%s %s default %s,", splice[0], splice[1], splice[2])
					native_parameters = append(native_parameters, splice[0])

				} else {

					if !is_jsonb {

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

			if in_declare {
				constantParameterName := regexp.MustCompile(`^[A-Z_]{1,}`).FindString(line_input)

				if constantParameterName != "" {
					debug_variables = append(debug_variables, constantParameterName)
					in_new_declare_statement = false
				} else {
					declare_splice := strings.SplitN(strings.TrimRight(line_input, ";"), " ", 3)
					splice := declare_splice[0]
					if in_new_declare_statement && splice != "" && len(declare_splice) >= 2 && declare_splice[1] != "record" {
						debug_variables = append(debug_variables, splice)
						in_new_declare_statement = false
					}
				}
				if strings.HasSuffix(strings.TrimSpace(line_input), ";") {
					in_new_declare_statement = true
				}

				if constantParameterName != "" && strings.Index(line_input, `:=`) == -1 {
					constantParameterType := strings.Trim(regexp.MustCompile(` [A-Za-z]{1,}`).FindString(line_input), " ")
					b.WriteString(fmt.Sprintf("%s constant %s := get_system_value_%s('%s');\n", constantParameterName, constantParameterType, constantParameterType, strings.ToLower(constantParameterName)))
					continue
				}
			}

			/* function name, passed parameters, and declared variables - done, now processing the rest */

			line_bytes := append([]byte(line_input), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}...)

			i := 0

			for line_bytes[i] > 0 {

				if string(line_bytes[i:i+7]) == `$BUILD$` {
					var build_version = BUILD_VERSION
					if is_debuggable {
						build_version += "*"
					}
					b.WriteString(build_version)
					i += 7
				}

				tag_processed = true

				for tag_processed {

					tag_processed = false /* to cover 2+ consecutive tags */

					/* comments */

					if (in_code || in_equals) && line_bytes[i] == '/' && line_bytes[i+1] == '*' {
						in_sql_comment = true
						i += 2
					}

					if (in_code || in_equals) && in_sql_comment && line_bytes[i] == '*' && line_bytes[i+1] == '/' {
						in_sql_comment = false
						i += 2
					}

					if !in_code && !in_equals {

						if string(line_bytes[i:i+4]) == `<!--` {
							in_html_comment = true
							i += 4
						}

						if string(line_bytes[i:i+3]) == `-->` && in_html_comment {
							in_html_comment = false
							i += 3
						}

						if string(line_bytes[i:i+5]) == "<form" {
							in_form = true
						}

						if line_bytes[i] == '<' {
							tag := re_tag.FindString(string(line_bytes[i+1:]))
							if tag != "" && !(tags_to_ignore[tag] || tags_to_ignore["/"+tag]) {
								tags_length := len(tags)
								if strings.HasPrefix(tag, "/") {
									if tags_length == 0 {
										fmt.Fprintf(os.Stderr, "\x1b[31;1mCannot close un-opened tag <%v> \x1b[0m \nFile: %v \nLine: %v \nTags: %v \n\n",
											tag, function_name, line_number, tags)
										os.Exit(1)
									}
									if "/"+tags[tags_length-1] != tag {
										fmt.Fprintf(os.Stderr, "\x1b[31;1mClosing tag <%v> mismatches opened tag <%v> \x1b[0m \nFile: %v \nLine: %v \nTags: %v \n\n",
											tag, tags[tags_length-1], function_name, line_number, tags)
										os.Exit(1)
									}
									tags = tags[:tags_length-1]
								} else {
									tags = append(tags, tag)
								}
								//fmt.Fprintf(os.Stderr, "tags at line %v: %v\n", line_number, tags)
							}
						}

					}

					/* quotes */

					if in_code || in_equals || in_declare {

						if !in_quotes_single && !in_quotes_backtick && line_bytes[i] == '"' {
							in_quotes_double = !in_quotes_double
						}

						if !in_quotes_double && !in_quotes_backtick && line_bytes[i] == '\'' {
							in_quotes_single = !in_quotes_single
						}

						if !in_quotes_single && !in_quotes_double && line_bytes[i] == '`' {
							in_quotes_backtick = !in_quotes_backtick
						}
					}

					if (in_code || in_equals || in_declare) && !(in_quotes_single || in_quotes_double || in_quotes_backtick) {

						if line_bytes[i] == '{' {
							in_jsonb = true
						}

						if line_bytes[i] == '!' && line_bytes[i+1] != '=' && line_bytes[i+1] != '~' {
							in_coalesce = true
							b.WriteString(`devnull(`)
							i++
							tag_processed = true
						}

						if lb := line_bytes[i]; in_coalesce {
							if !((lb >= 'a' && lb <= 'z') || (lb >= '0' && lb <= '9') || lb == '_' || lb == '.') {
								in_coalesce = false
								b.WriteString(`)`)
							}
						}

						if line_bytes[i] == '.' && line_bytes[i+1] == '.' && line_bytes[i+2] == '.' && line_bytes[i+3] == '{' {
							b.WriteString(`p_jsonb||`)
							i += 3
							tag_processed = true
						}

					}

					/* jsonb */

					if in_jsonb /*&& !(in_quotes_single || in_quotes_double || in_quotes_backtick)*/ {

						is_building_jsonb := true

						if line_bytes[i] == '\\' && line_bytes[i+1] == '{' {
							b.WriteString(`{`)
							i += 1
						} else if line_bytes[i] == '\\' && line_bytes[i+1] == '}' {
							b.WriteString(`}`)
							i += 1
						} else {

							switch line_bytes[i] {
							case '{':
								b.WriteString(`jsonb_build_object(`)
								jsonb++
							case '}':
								b.WriteString(`)`)
								jsonb--
								if jsonb == 0 {
									in_jsonb = false
								}
							case '[':
								b.WriteString(`jsonb_build_array(`)
							case ']':
								b.WriteString(`)`)
							case ':':
								if line_bytes[i-1] != ':' && line_bytes[i+1] != ':' && line_bytes[i+1] != '=' {
									b.WriteString(`,`)
								} else {
									is_building_jsonb = false
								}
							case '"':
								b.WriteString(`'`)
							default:
								is_building_jsonb = false
							}
						}
						if is_building_jsonb {
							i++
							tag_processed = true

						}
					}

					/* code tag <% %> but not print tag <%= %> */

					if !in_code && line_bytes[i] == '<' && line_bytes[i+1] == '%' && line_bytes[i+2] != '=' {
						in_code = true
						if !is_raw && !is_view {
							b.WriteString(fmt.Sprintf("'; _n_ := _n_ + 1; _l_ := %v; ", line_number))

							writePogoBreakPoint(&b, debug_variables, line_number)

							if is_pragma_timing {
								b.WriteString("_tm_ := _tm_ || jsonb_build_object('Line ' || _l_ || ':', ((clock_timestamp()-_tc_)::interval)); _tc_ := clock_timestamp(); ")
							}
						}

						i += 2
						tag_processed = true
					}

					if line_bytes[i] == '%' && line_bytes[i+1] == '>' {

						if in_code {
							in_code = false
							if !is_raw && !is_view {
								b.WriteString(`_v_[_n_] := '`)
							}
							i += 2
							tag_processed = true
						}

						if in_equals {
							in_equals = false
							b.WriteString(`)) || '`)
							i += 2
							tag_processed = true
						}
					}

					/* new style print tag <= => */
					if !in_code && !in_equals && line_bytes[i] == '<' && line_bytes[i+1] == '=' {
						in_equals = true
						b.WriteString(`' || (`)
						i += 2
						tag_processed = true
					}

					/* classic style print tag <%= %> */
					if !in_code && line_bytes[i] == '<' && line_bytes[i+1] == '%' && line_bytes[i+2] == '=' {
						in_equals = true
						b.WriteString(`' || coalesce2('' || (`)
						i += 3
						tag_processed = true
					}

					if !in_code && !in_equals {

						if in_form && line_bytes[i] == '<' {
							form++
						}

						if in_form && line_bytes[i] == '>' {
							form--
							// if form == 0 {
							// 	b.WriteByte(line_bytes[i])
							// 	line_bytes = append([]byte(`<input type="hidden" name="new_session_token" value="<%= f_get_session_token(CU) %>">`), line_bytes[i+1:]...)
							// 	i = 0
							// 	in_form = false
							// }
						}
					}

					if (in_code || in_equals) && line_bytes[i] == '`' {
						b.WriteString(`$tick$`)
						i++
					}

				} /* for tag_processed */

				/* removing comments, both pgsql and html */
				to_skip = false

				if (in_code || in_equals) && in_sql_comment {
					to_skip = true
				}

				if (!in_code && !in_equals) && in_html_comment {
					to_skip = true
				}

				if !to_skip {

					if !in_code && !in_equals && !in_declare && !in_header && line_bytes[i] == '\'' {
						b.WriteByte(line_bytes[i])
					}

					if line_bytes[i] > 0 {
						b.WriteByte(line_bytes[i])
					}

					if in_code && !is_view && !in_quotes_backtick && !in_quotes_double && !in_quotes_single && !in_jsonb && line_bytes[i] == ';' {
						writePogoBreakPoint(&b, debug_variables, line_number)
					}
				}

				i++

			} /* regular line */

		} /* not first line */

		b.WriteString("\n")

	} /* for scanner.scan */

	if len(tags) > 0 {
		fmt.Fprintf(os.Stderr, "\x1b[31;1mOpened tags not closed \x1b[0m \nFile: %v \nTags: %v \n\n",
			function_name, tags)
		os.Exit(1)
	}

	if !is_raw && !is_view {
		b.WriteString("';\n_n_ := _n_ + 1;\n")
	}

	b.WriteString(importFile("templates/template_psp_function_end" + template_suffix + ".sql"))

	fmt.Println(b.String())

	os.Exit(0)

} /* main */
