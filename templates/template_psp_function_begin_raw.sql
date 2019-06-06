\t
create function $function$ ($parameters$)
returns $returns$ as $$
declare
	_ds_ jsonb := jsonb_build_object();
	_ds_bs_ jsonb := jsonb_build_object();
    _frame_depth_ integer;
	_thread_id_ varchar := (coalesce(nullif(current_setting('__pogo.thread_id', true), ''), '<n/a>'));
