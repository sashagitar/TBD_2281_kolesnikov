# Задание 2
 
 1. [Однотабличные запросы](#1!!)
 2.
 3.
 4.

<br></br>

## <a name="1!!"></a>Однотабличные запросы

1. [Вывести всеми возможными способами имена и фамилии студентов, средний балл которых от 4 до 4.5](#1_0)
2. [Познакомиться с функцией CAST. Вывести при помощи неё студентов заданного курса (использовать Like)](#2_0)
3. [Вывести всех студентов, отсортировать по убыванию номера группы и имени от а до я](#3_0)
4. [Вывести студентов, средний балл которых больше 4 и отсортировать по баллу от большего к меньшему](#4_0)
5. [Вывести на экран название и риск футбола и хоккея](#5_0)
6. [Вывести id хобби и id студента которые начали заниматься хобби между двумя заданными датами (выбрать самим) и студенты должны до сих пор заниматься хобби](#6_0)
7. [Вывести студентов, средний балл которых больше 4.5 и отсортировать по баллу от большего к меньшему](#7_0)
8. [Из запроса №7 вывести несколькими способами на экран только 5 студентов с максимальным баллом](#8_0)
9. [Выведите хобби и с использованием условного оператора сделайте риск словами:](#9_0)
10. [Вывести 3 хобби с максимальным риском](#10_0)

###  <a name="1_0"></a> Вывести всеми возможными способами имена и фамилии студентов, средний балл которых от 4 до 4.5

#### `Запрос`

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where st.score <=4.5 and st.score >= 4
```

#### `Вывод`:

![exesize1](image/exesize1.png)

<br></br>

### <a name="2_0"></a> Познакомиться с функцией CAST. Вывести при помощи неё студентов заданного курса (использовать Like)

#### `Запрос`

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where cast(st.n_group as varchar) like '2281'

-- или

select st.name, st.surname, st.n_group, st.score
from students st
where st.n_group::varchar like '2281'
```

#### `Вывод`

![exesize2](image/exesize2.png)

<br></br>

### <a name="3_0"></a> Вывести всех студентов, отсортировать по убыванию номера группы и имени от а до я

#### `Запрос`

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
order by st.n_group desc, st.name 
```

#### `Вывод`

![exesize3](image/exesize3.png)

<br></br>

### <a name="4_0"></a> Вывести студентов, средний балл которых больше 4 и отсортировать по баллу от большего к меньшему

#### `Запрос`

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4
order by st.score desc
```

#### `Вывод`

![exesize4](image/exesize4.png)

<br></br>

### <a name="5_0"></a> Вывести на экран название и риск футбола и хоккея

#### `Запрос`

```SQL
select hb.name, hb.risk
from hobby hb
where name = 'футбол' or name = 'хоккей'
```

#### `Вывод`

![exesize5](image/exesize5.png)

<br></br>

### <a name="6_0"></a> Вывести id хобби и id студента которые начали заниматься хобби между двумя заданными датами (выбрать самим) и студенты должны до сих пор заниматься хобби

#### `Запрос`

```SQL
select st.student_id, st.hobby_id, st.date_start, st.date_finish
from students_hobbies st
where  date_start >= '2002-04-10' and date_finish is null;
```

#### `Вывод`

![exesize6](image/exesize6.png)

<br></br>

### <a name="7_0"></a> Вывести студентов, средний балл которых больше 4.5 и отсортировать по баллу от большего к меньшему

#### `Запрос`

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4.5
order by st.score 
```

#### `Вывод`

![exesize7](image/exesize7.png)

<br></br>

### <a name="8_0"></a> Из запроса №7 вывести несколькими способами на экран только 5 студентов с максимальным баллом

#### `Запрос`

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4.5
order by st.score 
/
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4.5 and surname like '%ич'
order by st.score 
/
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4.5 and n_group = 2281
order by st.score 
/
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4.5 and not n_group = 9003
order by st.score 
```

#### `Вывод`

![exesize8](image/exesize8_1.png)

![exesize8.2](image/exesize8_2.png)

![exesize8.3](image/exesize8_3.png)

![exesize8.4](image/exesize8_4.png)

<br></br>

### <a name="9_0"></a> Выведите хобби и с использованием условного оператора сделайте риск словами:

#### `Запрос`

```SQL
select hb.name, hb.risk,
case 
    when risk >=8 then 'очень высокий'
    when risk >= 6 and risk < 8 then 'высокий'
    when risk >= 4 and risk <8 then 'средний'
    when risk >= 2 and risk <4 then 'низкий'
    when risk < 2 then 'очень низкий'
end
from hobby hb
```

#### `Вывод`

![exesize8.4](image/exesize9.png)

<br></br>

### <a name="10_0"></a> Вывести 3 хобби с максимальным риском

#### `Запрос`

```SQL
select hb.name, hb.risk,
case 
    when risk >=8 then 'очень высокий'
    when risk >= 6 and risk < 8 then 'высокий'
    when risk >= 4 and risk <8 then 'средний'
    when risk >= 2 and risk <4 then 'низкий'
    when risk < 2 then 'очень низкий'
end
from hobby hb
```

#### `Вывод`

![exesize8.4](image/exesize10.png)

<br></br>

## Групповые функции

1. [Выведите на экран номера групп и количество студентов, обучающихся в них](#1_1)
2. [Выведите на экран для каждой группы максимальный средний балл](#1_2)
3. [Подсчитать количество студентов с каждой фамилией](#1_3)
4. [Подсчитать студентов, которые родились в каждом году](#1_4)
5. [Для студентов каждого курса подсчитать средний балл см. Substr](#1_5)
6. [Для студентов заданного курса вывести один номер группы с максимальным средним баллом](#1_6)
7. [Для каждой группы подсчитать средний балл, вывести на экран только те номера групп и их средний балл, в которых он менее или равен 3.5. Отсортировать по от меньшего среднего балла к большему.](#1_7)
8. [Для каждой группы в одном запросе вывести количество студентов, максимальный балл в группе, средний балл в группе, минимальный балл в группе](#1_8)
9. [Вывести студента/ов, который/ые имеют наибольший балл в заданной группе](#1_9)
10. [Аналогично 9 заданию, но вывести в одном запросе для каждой группы студента с максимальным баллом.](#1_10)

<br></br>

### <a name="1_1"></a> Выведите на экран номера групп и количество студентов, обучающихся в них

#### `Запрос`

```SQL
SELECT n_group,
       COUNT(n_group) AS stud_count
FROM students
GROUP BY n_group
ORDER BY n_group DESC;
```

#### `Вывод`

![exesize1_1](image/exe1_1.png)

<br></br>

### <a name="1_2"></a> Выведите на экран для каждой группы максимальный средний балл

#### `Запрос`

```SQL
SELECT n_group,
       max(score) AS score_max
FROM students
GROUP BY n_group
ORDER BY n_group DESC;
```

#### `Вывод`

![exesize1_2](image/exe1_2.png)

<br></br>

### <a name="1_3"></a> Подсчитать количество студентов с каждой фамилией

#### `Запрос`

```SQL
SELECT surname,
       COUNT(surname) AS surname_unik
FROM students
GROUP BY surname
ORDER BY surname DESC;
```

#### `Вывод`

![exesize1_3](image/exe1_3.png)

<br></br>

### <a name="1_4"></a> Подсчитать студентов, которые родились в каждом году

#### `Запрос`

```SQL
SELECT COUNT(extract(year from date_birth)) AS count_year, 
	extract(year from date_birth) as tg 
FROM students
GROUP BY tg
ORDER BY tg DESC;
```

#### `Вывод`

![exesize1_4](image/exe1_4.png)

<br></br>

### <a name="1_5"></a> Для студентов каждого курса подсчитать средний балл

#### `Запрос`

```SQL
SELECT left(n_group::varchar, 1) kurs,
	avg(score) as sr_score
FROM students
GROUP BY kurs
ORDER BY kurs DESC;
```

#### `Вывод`

![exesize1_5](image/exe1_5.png)

<br></br>

### <a name="1_6"></a> Для студентов заданного курса вывести один номер группы с максимальным средним баллом


#### `Запрос`

```SQL
SELECT max(score) as max_score, n_group
FROM students
where left(n_group::varchar, 1) = '2'
GROUP BY n_group
order by avg(score) desc
limit 1
```

#### `Вывод`

![exesize1_6](image/exe1_6.png)

<br></br>

### <a name="1_7"></a> Для каждой группы подсчитать средний балл, вывести на экран только те номера групп и их средний балл, в которых он менее или равен 3.5. Отсортировать по от меньшего среднего балла к большему.


#### `Запрос`

```SQL
SELECT avg(score) as max_score, n_group
FROM students
GROUP BY n_group
having avg(score) <= 3.5
order by avg(score) asc
```

#### `Вывод`

![exesize1_7](image/exe1_7.png)

<br></br>

### <a name="1_8"></a> Для каждой группы в одном запросе вывести количество студентов, максимальный балл в группе, средний балл в группе, минимальный балл в группе


#### `Запрос`

```SQL
SELECT n_group, COUNT(n_group) as num_stud, 
    max(score) as max_score, 
    avg(score) as avg_score, 
    min(score) as min_score
FROM students
GROUP BY n_group
```

#### `Вывод`

![exesize1_8](image/exe1_8.png)

<br></br>

### <a name="1_9"></a> Вывести студента/ов, который/ые имеют наибольший балл в заданной группе

#### `Запрос`

```SQL
select st.*
from(SELECT n_group, max(score) as max_score
	 FROM students
	 GROUP BY n_group
	 having n_group = 2281
) t_my, students st
where st.score = t_my.max_score
```

#### `Вывод`

![exesize1_9](image/exe1_9.png)

<br></br>

### <a name="1_10"></a> Аналогично 9 заданию, но вывести в одном запросе для каждой группы студента с максимальным баллом.

#### `Запрос`

```SQL
select st.*
from(SELECT n_group, max(score) as max_score
	 FROM students
	 GROUP BY n_group
) t_my, students st
where st.n_group = t_my.n_group and st.score = t_my.max_score
```

#### `Вывод`

![exesize1_10](image/exe1_10.png)