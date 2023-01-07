# plpgsql

## Задания

Используя только средства PL/pgSQL выполните следующие задания:

1. Выведите на экран любое сообщение

```sql
CREATE OR REPLACE FUNCTION msg() returns varchar
AS $$
BEGIN 
	return 'msg';
END
$$ language plpgsql;

SELECT msg();
```

2. Выведите на экран текущую дату

```sql
create or replace function date_today() returns varchar
as $$
begin
	return now();
end
$$ language plpgsql;

select date_today();
```


3. Создайте две числовые переменные и присвойте им значение. Выполните математические действия с этими числами и выведите результат на экран.

```sql
create or replace function math(OUT sum int, OUT multi int, OUT sub int, OUT div float) 
as $$
declare
x int;
y int;
begin
	x := 7;
	y := 5;
    sum = x + y;
    multi = x * y;
    sub = x - y;
    div = x/y;
	
end;
$$ language plpgsql;

select * from math();
```

4. Написать программу двумя способами 1 - использование IF, 2 - использование CASE. Объявите числовую переменную и присвоейте ей значение. Если число равно 5 - выведите на экран "Отлично". 4 - "Хорошо". 3 - Удовлетворительно". 2 - "Неуд". В остальных случаях выведите на экран сообщение, что введённая оценка не верна.

```sql
-- if
create or replace function with_if(n int) returns varchar
as $$
begin
	if n = 5 then return 'Отлично';
	elsif n = 4 then return 'Хорошо';
	elsif n = 3 then return 'Удовлетворительно';
	elsif n = 2 then return 'Неуд';
	else return 'введённая оценка не верна';
	end if;
	
end
$$ language plpgsql;

select with_if(5);
```

```sql
-- case
create or replace function with_case(n int) returns varchar
as $$
begin
	case n
		when 5 then return 'Отлично';
		when 4 then return 'Хорошо';
		when 3 then return 'Удовлетворительно';
		when 2 then return 'Неуд';
		else return 'введённая оценка не верна';
	end case;
end
$$ language plpgsql;

select with_case(2);
```


5. Выведите все квадраты чисел от 20 до 30 3-мя разными способами (LOOP, WHILE, FOR).

```sql
-- loop
create or replace procedure pow_numbers_loop()
language plpgsql
as $$
declare i int = 20;
begin
	loop
		exit when i > 30;
		raise notice 'number: %',i*i;
		i = i + 1;
	end loop;
end;
$$;

call pow_numbers_loop()
```

```sql
-- while
create or replace procedure pow_numbers_while()
language plpgsql
as $$
declare 
	i int := 20;
begin
	while i <= 30 loop
		raise notice 'number: %',i*i;
		i = i + 1;
	end loop;
end;
$$;

call pow_numbers_while()
```

```sql
create or replace procedure pow_numbers_for()
language plpgsql
as $$
begin
	for i in 20..30 loop
		raise notice 'number: %',i*i;
	end loop;
end;
$$;

call pow_numbers_for()
```

6. Последовательность Коллатца. Берётся любое натуральное число. Если чётное - делим его на 2, если нечётное, то умножаем его на 3 и прибавляем 1. Такие действия выполняются до тех пор, пока не будет получена единица. Гипотеза заключается в том, что какое бы начальное число n не было выбрано, всегда получится 1 на каком-то шаге. **Задания:** написать функцию, входной параметр - начальное число, на выходе - количество чисел, пока не получим 1; написать процедуру, которая выводит все числа последовательности. Входной параметр - начальное число.

```sql
create or replace function collatz(x int) returns int
as $$
begin
	while x!=1 loop
		case 
			when x%2=0 then x = x/2;
			else x = x*3+1;
		end case;
	end loop;
	return x;
		
end
$$ language plpgsql;

select collatz(3);
```

```sql
create or replace procedure collatz_p(x int)
language plpgsql
as $$
begin
	while x!=1 loop
		raise notice '%', x;
		case 
			when x%2=0 then x = x/2;
			else x = x*3+1;
		end case;
	end loop;
		
end
$$;

call collatz_p(7);
```


7. Числа Люка. Объявляем и присваиваем значение переменной - количество числе Люка.
Вывести на экран последовательность чисел. Где `L0 = 2, L1 = 1` ; `Ln=Ln-1 + Ln-2` (сумма двух предыдущих чисел). **Задания:** написать фунцию, входной параметр - количество чисел, на выходе - последнее число (Например: входной 5, 2 1 3 4 7 - на выходе число 7); написать процедуру, которая выводит все числа последовательности. Входной параметр - количество чисел.

```sql
create or replace function luka(n int) returns int
language plpgsql
as $$
declare
n1 int;
n2 int;
count_ int;
begin
	n1 := 2;
	n2 := 1;
	count_ := 2;
	while (n > count_) loop
		count_ := count_ + 1;
		n2 := n1 + n2;
		n1 := n2 - n1;
	end loop;
	return	n2;
end;
$$;

select luka(5)
```

```sql
create or replace procedure luka_p(n int)
language plpgsql
as $$
declare
n1 int;
n2 int;
count_ int;
begin
	n1 := 2;
	n2 := 1;
	count_ := 2;
	raise notice '2';
	raise notice '1';
	while (n > count_) loop
		count_ := count_ + 1;
		n2 := n1 + n2;
		n1 := n2 - n1;
		raise notice '%', n2;
	end loop;
end;
$$;

call luka_p(5)
```

8. Напишите функцию, которая возвращает количество человек родившихся в заданном году.

```sql
create or replace function brdin(year int) returns int
language plpgsql
as $$
declare
count_ int;
begin
	select count(*) into count_
	from people
	where extract(year from people.birth_date) = brdin.year;
	return count_;
end;
$$;

select brdin(1995)
```

9. Напишите функцию, которая возвращает количество человек с заданным цветом глаз.

```sql
create or replace function count_people_by_color_eyes(color varchar) returns int
language plpgsql
as $$
declare
count_ int;
begin
	select count(*) into count_
	from people
	where people.eyes = color;
	return count_;
end;
$$;

select count_people_by_color_eyes('brown')
```

10. Напишите функцию, которая возвращает ID самого молодого человека в таблице.

```sql
create or replace function yong() returns int
language plpgsql
as $$
declare
id_ int;
begin
	select id into id_
	from people
	order by people.birth_date desc
	limit 1;
	return id_;
end;
$$;

select yong()
```

11. Напишите процедуру, которая возвращает людей с индексом массы тела больше заданного. ИМТ = масса в кг / (рост в м)^2.

```sql
create or replace function mass(index_ int) returns table (id int)
language plpgsql
as $$
declare
id_ int;
begin
	return query 
		select people.id
		from people
		where people.weight / (people.growth/100)*(people.growth/100) > index_;
end;
$$;

select * from mass(50)
```

12. Измените схему БД так, чтобы в БД можно было хранить родственные связи между людьми. Код должен быть представлен в виде транзакции (Например (добавление атрибута): `BEGIN; ALTER TABLE people ADD COLUMN leg_size REAL; COMMIT;`). Дополните БД данными.

```sql
begin;
alter table people add column parent int;
commit;
```

13. Напишите процедуру, которая позволяет создать в БД нового человека с указанным родством.

```sql
create or replace procedure add_cilde(id_pr int, name_ varchar) 
language plpgsql
as $$
begin
	insert into people (name, parent) values (name_, id_pr);
end;
$$;

call add_cilde(3, 'lol')
```

14. Измените схему БД так, чтобы в БД можно было хранить время актуальности данных человека (выполнить также, как п.12).

```sql
begin;
alter table people add column last_update date;
update people set last_update=now();
commit;
```

15. Напишите процедуру, которая позволяет актуализировать рост и вес человека.

```sql
create or replace procedure update_data(id_ int, weight_ real, height_ real) 
language plpgsql
as $$
begin
	update people set growth=height_, weight=weight_, last_update=now() where id=id_;
end;
$$;

call update_data(2, 63.8, 190)
```

## Мои контакты
![vk](vk_icon.png)  https://vk.com/hidden_by.the_devil

Email: DAI.20@uni_dubna.ru
