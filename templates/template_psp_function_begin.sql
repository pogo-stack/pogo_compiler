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
	_h_ integer := 200;
	_errid_ varchar := '';
	_tc_ timestamp := clock_timestamp();
	_tm_ jsonb := '[]'::jsonb;
	_r_ pogo__return__type;
	v_error_stack text := '';
