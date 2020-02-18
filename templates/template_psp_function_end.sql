if _d_ > 0 then
	_v_[_n_] := '<!-- end: $name$ -->'; _n_ := _n_ + 1;
	_v_[_n_] := '<!-- ' || (clock_timestamp() - _c_) :: interval || ' -->'; _n_ := _n_ + 1;
	if _tm_ != '[]'::jsonb then
		_tm_ := _tm_ || jsonb_build_object('Line end:', ((clock_timestamp()-_tc_)::interval));
		_v_[_n_] := '<!-- ' || jsonb_pretty(_tm_) || ' -->'; _n_ := _n_ + 1;
	end if;
end if;
_c_ := clock_timestamp();


for i in 1 .. _n_ loop
	if _v_[i] is null then _v_[i] := ''; end if;
end loop;

_t_ := array_to_string(_v_, '');

return (row(_t_::text,
			_jsonr_::jsonb,
			_bytear_::bytea,
			_addr_::jsonb
			)
		);

exception when others then
	get stacked diagnostics v_error_stack = pg_exception_context;

	_h_ := 0;

	loop
		_h_ := _h_ + 1;
		exit when _h_ > 2;
		begin
			insert into __pogo__errors
		
	(
				id,
				time_stamp,
				user_id,
				function_name,
				params,
				line_number,
				stack,
				sql_state,
				sql_error
			)
			values
			(
				md5(clock_timestamp()::varchar || _l_),
				current_timestamp,
				1,
				'$function$',
				$form$,
				_l_,
				v_error_stack::varchar,
				sqlstate,
				sqlerrm
			)
			returning id into _errid_;
			raise notice 'sql error: % - % - %', '$function$'::text, sqlerrm::text, v_error_stack::text;
			exit;
			exception when others then null;
		end;
	end loop;

	return (row(
					'<pre class="pre-yellow-error" style="background-color:yellow; color:pink;">*** error has occurred '
								|| (case when 1 in (1) then '*** <a target="_blank" href="error?p_id=' || _errid_ || '">details</a>' else '' end)
								|| ' ***</pre>'::text,
					jsonb_build_object(
								'status', 'error',
								'error_id', _errid_
							)::jsonb,
					null::bytea,
					jsonb_build_object('http_code', 500)
				)
			);


end;
$$
language plpgsql $volatility_category$;

\t

delete from __pogo__testcases where page_name = '$name$';

insert into __pogo__testcases
select distinct '$name$', unnest(regexp_split_to_array($testcases$, ','));

delete from __pogo__dependencies where page_name = '$name$';

insert into __pogo__dependencies
select distinct '$name$', unnest(regexp_split_to_array($dependencies$, ','));

do $$
begin
	insert into __pogo__compiled_source (name, latest_version)
	values ('$name$', 1);
exception when others then
	update __pogo__compiled_source
	set latest_version = latest_version + 1
	where name = '$name$';
end;
$$;

update __pogo__compiled_source set is_noauth = $is_noauth$ where name = '$name$';


$debugger_breakpoints$
