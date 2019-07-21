\t
create function $function$ ($parameters$)
returns $returns$ as $$
declare
	_c_ timestamp := clock_timestamp();
	_t_ text;
	_l_ integer := 0;
	_v_ varchar[];
	_n_ integer := 1;
	_d_ integer := $debug$;
	_ds_ jsonb := jsonb_build_object();
	_ds_bs_ jsonb := jsonb_build_object();
	_frame_depth_ integer;
	_thread_id_ varchar := (coalesce(nullif(current_setting('__pogo.thread_id', true), ''), '<n/a>'));
	_errid_ varchar := '';
	_tc_ timestamp := clock_timestamp();
	_tm_ jsonb := jsonb_build_array();
	_h_ integer := 0;
	v_error_stack text := '';
	_jsonr_ jsonb := jsonb_build_object();
	_bytear_ bytea := null::bytea;
	_addr_ jsonb := jsonb_build_object('http_code', 200, 'headers', jsonb_build_array(), 'cookies', jsonb_build_array());