# Задание 2
 
 1. [Однотабличные запросы](#1!!)
 2. [Групповые функции](#2!!)
 3. [Многотабличные запросы](#3!!)

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

## <a name="2!!"></a>Групповые функции

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

<br></br>

## <a name="3!!"></a>Многотабличные запросы

1. [Вывести все имена и фамилии студентов, и название хобби, которым занимается этот студент.](#3_1)
2. [Вывести информацию о студенте, занимающимся хобби самое продолжительное время.](#3_2)
3. [Вывести имя, фамилию, номер зачетки и дату рождения для студентов, средний балл которых выше среднего, а сумма риска всех хобби, которыми он занимается в данный момент, больше 0.9.](#3_3)
4. [Вывести фамилию, имя, зачетку, дату рождения, название хобби и длительность в месяцах, для всех завершенных хобби Диапазон дат.](#3_4)
5. [Вывести фамилию, имя, зачетку, дату рождения студентов, которым исполнилось N полных лет на текущую дату, и которые имеют более 1 действующего хобби.](#3_5)
6. [Найти средний балл в каждой группе, учитывая только баллы студентов, которые имеют хотя бы одно действующее хобби.](#3_6)
7. [Найти название, риск, длительность в месяцах самого продолжительного хобби из действующих, указав номер зачетки студента.](#3_7)
8. [Найти все хобби, которыми увлекаются студенты, имеющие максимальный балл.](#3_8)
9. [Найти все действующие хобби, которыми увлекаются троечники 2-го курса.](#3_9)
10. [Найти номера курсов, на которых более 50% студентов имеют более одного действующего хобби.](#3_10)
11. [Вывести номера групп, в которых не менее 60% студентов имеют балл не ниже 4.](#3_11)
12. [Для каждого курса подсчитать количество различных действующих хобби на курсе.](#3_12)
13. [Вывести номер зачётки, фамилию и имя, дату рождения и номер курса для всех отличников, не имеющих хобби. Отсортировать данные по возрастанию в пределах курса по убыванию даты рождения.](#3_13)
14. [Создать представление, в котором отображается вся информация о студентах, которые продолжают заниматься хобби в данный момент и занимаются им как минимум 5 лет.](#3_14)
15. [Для каждого хобби вывести количество людей, которые им занимаются.](#3_15)
16. [Вывести ИД самого популярного хобби.](#3_16)
17. [Вывести всю информацию о студентах, занимающихся самым популярным хобби.](#3_17)
18. [Вывести ИД 3х хобби с максимальным риском.](#3_18)
19. [Вывести 10 студентов, которые занимаются одним (или несколькими) хобби самое продолжительно время.](#3_19)
20. [Вывести номера групп (без повторений), в которых учатся студенты из предыдущего запроса.](#3_20)
21. [Создать представление, которое выводит номер зачетки, имя и фамилию студентов, отсортированных по убыванию среднего балла.](#3_21)
22. [Представление: найти каждое популярное хобби на каждом курсе.](#3_22)
23. [Представление: найти хобби с максимальным риском среди самых распространенных хобби на 2 уровне.](#3_23)
24. [Представление: для каждого курса подбирается количество студентов на курсе и количество отличников.](#3_24)
25. [Представление: самое популярное хобби среди всех студентов.](#3_25)
26. [Создать обновляемое представление.](#3_26)
27. [Для каждой буквы алфавита из числа найденных величин, средних и балльных. (Т.е. среди всех студентов, чьё имя начинается на А (Алексей, Алина, Артур, Анджела) найти то, что указано в задании. Вывести на экран тех, средний балл проходит больше 3,6](#3_27)
28. [Для каждой фамилии характерна высокая степень тяжести и снижение балла. (Например, в университете учатся 4 Иванова (1-2-3-4). 1-2-3 учатся на 2 курсах и имеют средний балл 4.1, 4, 3.8, соответственно, а 4 Иванов учится на 3 курсах и имеет балл 4.5. На экране должно быть следующее:](#3_28)

id | surname | max | min
---|---------|-----|----    
2  | Иванов  | 4,1 | 3,8
3  | Иванов  | 4,5 | 4,5

29. [Для каждого года рождения подсчитывается количество хобби, наблюдаемое или занимавшееся студентами.](#3_29)
30. [Для каждой буквы алфавита в большинстве случаев возникает риск, связанный с хобби.](#3_30)
31. [Для каждого месяца из-за даты рождения студентов получаются средние баллы, которые являются хобби под названием «Футбол».](#3_31)
32. [Вывести информацию о студентах, которые занимались или занимались хотя бы 1 хобби в следующем формате: Имя: Иван, фамилия: Иванов, группа: 1234](#3_32)
33. [Внешний вид в какой-то по счёту символа встречается «ов». Если 0 (т.е. не встречается, то выведите на экран «не найдено»).](#3_33)
34. [Дополнить фамилию прямым символом # до 10 символов.](#3_34)
35. [При помощи функции удалить все символы # из полученного запроса.](#3_35)
36. [Выведите на экран количество дней в прошлом 2018 году.](#3_36)
37. [Вывести на экран какого числа будет ближайшая суббота.](#3_37)
38. [Выведите на экран век, а также какая сейчас неделя года и день года.](#3_38)
39. [Вы вводите всех студентов, которые занимались хотя бы 1 хобби. Выведите на экран Имя, Фамилию, Название хобби, а также надпись «занимается», если студент продолжает заниматься хобби в данный момент или «закончил», если уже не занимается.](#3_39)
40. [Для каждой группы приводят количество студентов учится на 5,4,3,2. Использовать обычное математическое округление. Итоговый результат должен выглядеть примерно так:](#3_40)

    | СЧЕТ | 2222 | 3011 | 4011 | 4032 |
    | ----- | ---- | ---- | ---- | ---- |
    | 2 | 0 | 0 | 0 | 1 |
    | 3 | 1 | 2 | 1 | 1 |
    | 4 | 4 | 3 | 3 | 3 |
    | 5 | 1 | 1 | 1 | 0 |

<br></br>

### <a name="3_1"></a> 1. Вывести все имена и фамилии студентов, и название хобби, которым занимается этот студент.

#### `Запрос`

```SQL

SELECT students.id,
       students.name,
       students.surname,
       hobby.name
FROM students,
     hobby, 
     students_hobbies
WHERE   hobby.id= students_hobbies.hobby_id and 
        students.id= students_hobbies.student_id and 
        students_hobbies.date_start is not null;
```

#### `Вывод`

![exesize3_1](image/exe3_1.png)

<br></br>

### <a name="3_2"></a> 2. Вывести информацию о студенте, занимающимся хобби самое продолжительное время.

#### `Запрос`

```SQL
SELECT 	st.id, st.name, st.surname, hb.name,
		st_hb.date_finish - st_hb.date_start days
FROM students st,
     hobby hb, 
	 students_hobbies st_hb
where   hb.id= st_hb.hobby_id and 
        st.id= st_hb.student_id and 
        st_hb.date_finish - st_hb.date_start is not null
order by days DESC
limit 1
```

#### `Вывод`

![exesize3_2](image/exe3_2.png)

<br></br>

### <a name="3_3"></a> 3. Вывести имя, фамилию, номер зачетки и дату рождения для студентов, средний балл которых выше среднего, а сумма риска всех хобби, которыми он занимается в данный момент, больше 9.

#### `Запрос`

```SQL
SELECT 	st.name, st.surname, 
        st.n_group,st.date_birth, 
        hb.name

FROM 	students st,
     	hobby hb, 
	 	students_hobbies st_hb
where hb.id= st_hb.hobby_id and 
      st.id= st_hb.student_id and 
      st_hb.date_finish is null and
	  st.score > (SELECT AVG(st.score) FROM students st) and
	  hb.risk > 0.9

```

#### `Вывод`

![exesize3_3](image/exe3_3.png)

<br></br>

### <a name="3_4"></a> 4. Вывести фамилию, имя, зачетку, дату рождения, название хобби и длительность в месяцах, для всех завершенных хобби Диапазон дат.

#### `Запрос`

```SQL
SELECT 	st.name, st.surname, st.n_group,st.date_birth, hb.name, (st_hb.date_finish - st_hb.date_start)/30 as month
FROM 	students st,
     	hobby hb, 
	 	students_hobbies st_hb
where 	hb.id= st_hb.hobby_id and 
      	st.id= st_hb.student_id and 
		st_hb.date_finish is not null
```

#### `Вывод`

![exesize3_4](image/exe3_4.png)

<br></br>

### <a name="3_5"></a> 5. Вывести фамилию, имя, зачетку, дату рождения студентов, которым исполнилось N полных лет на текущую дату, и которые имеют более 1 действующего хобби.

#### `Запрос`

```SQL
SELECT 	st.id, st.name, st.surname, st.n_group, 
        st.date_birth, COUNT(st_hb.hobby_id) count
FROM 	students st,
     	hobby hb, 
	 	students_hobbies st_hb
where 	hb.id= st_hb.hobby_id and 
      	st.id= st_hb.student_id and 
		date_part('year', now()) - date_part('year', st.date_birth) > 19
		GROUP BY st.id,  st.name, st.surname, st.n_group, st.date_birth
		HAVING COUNT(st_hb.hobby_id) > 1
```

#### `Вывод`

![exesize3_5](image/exe3_5.png)

<br></br>

### <a name="3_6"></a> 6. Найти средний балл в каждой группе, учитывая только баллы студентов, которые имеют хотя бы одно действующее хобби.

#### `Запрос`

```SQL
SELECT 	st.n_group, avg(st.score) sr
FROM 	students st,
     	hobby hb, 
	 	students_hobbies st_hb
where 	hb.id= st_hb.hobby_id and 
      	st.id= st_hb.student_id
GROUP BY st.n_group
```

#### `Вывод`

![exesize3_6](image/exe3_6.png)

<br></br>

### <a name="3_7"></a> 7. Найти название, риск, длительность в месяцах самого продолжительного хобби из действующих, указав номер зачетки студента.

#### `Запрос`

```SQL
SELECT 	hb.name, hb.risk, max(st_hb.date_finish - st_hb.date_start)/30 max_time, st.name, st.surname
FROM 	students st,
     	hobby hb, 
	 	students_hobbies st_hb
where 	hb.id= st_hb.hobby_id and 
      	st.id= st_hb.student_id
GROUP BY hb.name, hb.risk, st.name, st.surname
HAVING max(st_hb.date_finish - st_hb.date_start)/30 is not null
order by max_time DESC
limit 1
```

#### `Вывод`

![exesize3_7](image/exe3_7.png)

<br></br>

### <a name="3_8"></a> 8. Найти все хобби, которыми увлекаются студенты, имеющие максимальный балл.

#### `Запрос`

```SQL
SELECT 	hb.name, max(st.score) max_score
FROM 	students st,
     	hobby hb, 
	 	students_hobbies st_hb
where 	hb.id= st_hb.hobby_id and 
      	st.id= st_hb.student_id
GROUP BY hb.name
```

#### `Вывод`

![exesize3_8](image/exe3_8.png)

<br></br>

### <a name="3_9"></a> 9. Найти все действующие хобби, которыми увлекаются троечники 2-го курса.

#### `Запрос`

```SQL
SELECT 	hb.name, st.name, st.surname, st.score
FROM 	students st,
     	hobby hb, 
	 	students_hobbies st_hb
where 	hb.id= st_hb.hobby_id and 
      	st.id= st_hb.student_id and
		st_hb.date_finish is null and
		st.n_group >1999 and
		st.n_group < 3000 and
		st.score <= 4.51
```

#### `Вывод`

![exesize3_9](image/exe3_9.png)

<br></br>

### <a name="3_10"></a> 10. Найти номера курсов, на которых более 50% студентов имеют более одного действующего хобби.

#### `Запрос`

```SQL
SELECT *
FROM
  (SELECT LEFT(st.n_group::VARCHAR,1), COUNT(st.id) 
    FROM students st
    GROUP BY
      LEFT(st.n_group::VARCHAR,1)) total
INNER JOIN
  (SELECT LEFT(st.n_group::VARCHAR,1), COUNT(st.id)
    FROM
      students st,
      (SELECT st.id, COUNT(st.id) FROM students st
        INNER JOIN students_hobbies sth
        ON sth.id = st.id
        INNER JOIN hobby h
        ON sth.hobby_id = h.id
        WHERE sth.date_finish IS NULL
        GROUP BY st.id
        HAVING COUNT(st.id) > 1) morethanone
    WHERE
      st.id = morethanone.id
    GROUP BY
      LEFT(st.n_group::VARCHAR,1)) morethanone
ON total.left = morethanone.left
WHERE
  total.count / 2 < morethanone.count
```

`взял у Кости, слишком уж геморное задание`

#### `Вывод`

![exesize3_10](image/exe3_10.png)

<br></br>

### <a name="3_11"></a> 11. Вывести номера групп, в которых не менее 60% студентов имеют балл не ниже 4.

#### `Запрос`

```SQL
SELECT sub.n_group
FROM
  (SELECT 
    st.n_group, 
    COUNT(st.id) total_count, 
    COUNT(st.score) FILTER (WHERE st.score > 4) above_score_count
  FROM students st
  GROUP BY st.n_group) sub
WHERE sub.total_count*0.6 < above_score_count
```

#### `Вывод`

![exesize3_11](image/exe3_11.png)

<br></br>

### <a name="3_12"></a> 12.  Для каждого курса подсчитать количество различных действующих хобби на курсе.

#### `Запрос`

```SQL
select left(st.n_group::varchar, 1) cours, count(hb.name) col_vo_hobby
from students st, students_hobbies st_hb, hobby hb
where st.id = st_hb.student_id and
		hb.id = st_hb.hobby_id
group by cours
```

#### `Вывод`

![exesize3_12](image/exe3_12.png)

<br></br>

### <a name="3_13"></a>  13. Вывести номер зачётки, фамилию и имя, дату рождения и номер курса для всех отличников, не имеющих хобби. Отсортировать данные по возрастанию в пределах курса по убыванию даты рождения.

#### `Запрос`

```SQL
select st.id, st.name, st.surname, st.date_birth, left(n_group::varchar, 1) n_cours
from hobby hb, students st, students_hobbies st_hb
where (st.id not in -- кто никогда не имел хобби
			(select st_hb.student_id
			from students_hobbies st_hb
			group by st_hb.student_id) or
	   (st.id in    -- кто имел но все закончил
			(SELECT sth.id
    		FROM students_hobbies sth
    		GROUP BY sth.id
    		HAVING COUNT(sth.date_finish) = COUNT(sth.date_start)))) and
	   st.score >= 4.5
group by st.id, n_cours,  st.name, st.surname, st.date_birth
```

#### `Вывод`

![exesize3_13](image/exe3_13.png)

<br></br>

### <a name="3_14"></a> 14. Создать представление, в котором отображается вся информация о студентах, которые продолжают заниматься хобби в данный момент и занимаются им как минимум 5 лет.

#### `Запрос`

```SQL
select st.id, st.name, st.surname, (now()::date - st_hb.date_start)/365 years
from hobby hb, students st, students_hobbies st_hb
where 	st_hb.student_id = st.id and
		st_hb.hobby_id = hb.id and
		st_hb.date_finish is null and
		(now()::date - st_hb.date_start)/365 > 5
		
group by st.id,  st.name, st.surname, st_hb.date_start
```

#### `Вывод`

![exesize3_14](image/exe3_14.png)

<br></br>

### <a name="3_15"></a> 15. Для каждого хобби вывести количество людей, которые им занимаются.

#### `Запрос`

```SQL
SELECT h.name, COUNT(DISTINCT sth.id)
FROM hobby h
INNER JOIN students_hobbies sth
ON h.id = sth.hobby_id
GROUP BY h.name
```

#### `Вывод`

![exesize3_15](image/exe3_15.png)

<br></br>

### <a name="3_16"></a> 16. Вывести ИД самого популярного хобби.

#### `Запрос`

```SQL
SELECT h.name, COUNT(DISTINCT sth.id) counte
FROM hobby h
INNER JOIN students_hobbies sth
ON h.id = sth.hobby_id
GROUP BY h.name
order by counte desc
limit 1
```

#### `Вывод`

![exesize3_16](image/exe3_16.png)

<br></br>

### <a name="3_17"></a> 17. Вывести всю информацию о студентах, занимающихся самым популярным хобби.

#### `Запрос`

```SQL
select st.id, st.name, st.surname, st.n_group
from students st
INNER JOIN students_hobbies st_hb
on st.id = st_hb.student_id
where 
	st_hb.hobby_id = (SELECT st_hb.hobby_id
    	FROM students_hobbies st_hb
    	GROUP BY st_hb.hobby_id
    	ORDER BY COUNT(st_hb.id) DESC
    	LIMIT 1) and
	st_hb.date_finish is null

```

#### `Вывод`

![exesize3_17](image/exe3_17.png)

<br></br>

### <a name="3_18"></a> 18. Вывести ИД 3х хобби с максимальным риском.

#### `Запрос`

```SQL
select hb.id, hb.name, hb.risk
from hobby hb
order by hb.risk desc
limit 3
```

#### `Вывод`

![exesize3_18](image/exe3_18.png)

<br></br>

### <a name="3_19"></a> 19. Вывести 10 студентов, которые занимаются одним (или несколькими) хобби самое продолжительно время.

#### `Запрос`

```SQL
select st.id, st.name, st.surname, hb.name, now()::date - st_hb.date_start days
from hobby hb, students st, students_hobbies st_hb
where 	st.id = st_hb.student_id and
		hb.id = st_hb.hobby_id and
		st_hb.date_finish is null
order by days desc
limit 10
```

#### `Вывод`

![exesize3_19](image/exe3_19.png)

<br></br>

### <a name="3_20"></a> 20. Вывести номера групп (без повторений), в которых учатся студенты из предыдущего запроса.

#### `Запрос`

```SQL
select tb.n_group
from (
	select st.id, st.n_group, st.name, st.surname, hb.name, now()::date - st_hb.date_start days
	from hobby hb, students st, students_hobbies st_hb
	where 	st.id = st_hb.student_id and
			hb.id = st_hb.hobby_id and
			st_hb.date_finish is null
	order by days desc
	limit 10) tb
group by tb.n_group
```

#### `Вывод`

![exesize3_20](image/exe3_20.png)

<br></br>

### <a name="3_21"></a> 21. Создать представление, которое выводит номер зачетки, имя и фамилию студентов, отсортированных по убыванию среднего балла.

#### `Запрос`

```SQL
select st.id, st.n_group, st.name, st.surname, st.score
from students st
order by st.score desc
```

#### `Вывод`

![exesize3_21](image/exe3_21.png)

<br></br>

### <a name="3_22"></a> 22. Представление: найти каждое популярное хобби на каждом курсе.

#### `Запрос`

```SQL
CREATE OR REPLACE VIEW h_mostpopular AS

select DISTINCT ON (1) left(st.n_group::varchar, 1) n_cours, hb.id

from students st
INNER JOIN students_hobbies st_hb
ON st.id = st_hb.student_id
INNER JOIN hobby hb
ON st_hb.hobby_id = hb.id

group by st.n_group, hb.id
order by n_cours, hb.id;

SELECT * FROM public.h_mostpopular
```

#### `Вывод`

![exesize3_22](image/exe3_22.png)

<br></br>

### <a name="3_23"></a> 23. Представление: найти хобби с максимальным риском среди самых распространенных хобби на 2 уровне.

#### `Запрос`

```SQL
CREATE OR REPLACE VIEW h_mostrisk_2grade AS
SELECT *
FROM hobby hb
WHERE hb.id
IN
  (SELECT hb.id
  FROM students st
  INNER JOIN students_hobbies st_hb
  ON st.id = st_hb.student_id
  INNER JOIN hobby hb
  ON st_hb.hobby_id = hb.id
  WHERE LEFT(st.id::VARCHAR,1) = '2'
  GROUP BY hb.id
  ORDER BY COUNT(hb.id))
ORDER BY hb.risk DESC
LIMIT 1;

SELECT * FROM public.h_mostrisk_2grade
```

#### `Вывод`

![exesize3_23](image/exe3_23.png)

<br></br>

### <a name="3_24"></a> 24. Представление: для каждого курса подбирается количество студентов на курсе и количество отличников.

#### `Запрос`

```SQL
CREATE OR REPLACE VIEW st_idealscore_bygrade AS
SELECT 
  LEFT(st.n_group::VARCHAR,1) kours, 
  COUNT(st.id) total, 
  COUNT(st.id) FILTER (WHERE st.score >= 4.5) goodcount
FROM students st
GROUP BY LEFT(st.n_group::VARCHAR,1);

SELECT * FROM public.st_idealscore_bygrade
```

#### `Вывод`

![exesize3_24](image/exe3_24.png)

<br></br>

### <a name="3_25"></a> 25. Представление: самое популярное хобби среди всех студентов.

#### `Запрос`

```SQL
CREATE OR REPLACE VIEW h_popular AS
SELECT *
FROM hobby hb
WHERE 
  hb.id = 
    (SELECT hb.id
    FROM students st
    INNER JOIN students_hobbies st_hb
    ON st.id = st_hb.student_id
    INNER JOIN hobby hb
    ON st_hb.hobby_id = hb.id
    GROUP BY hb.id
    ORDER BY COUNT(hb.id) DESC
    LIMIT 1);
	
SELECT * FROM public.h_popular
```

#### `Вывод`

![exesize3_25](image/exe3_25.png)

<br></br>

### <a name="3_26"></a> 26. Создать обновляемое представление.

#### `Запрос`

```SQL
CREATE OR REPLACE VIEW students_short AS
SELECT st.id, st.name, st.surname, st.n_group
FROM students st;
SELECT * FROM public.students_short
```

#### `Вывод`

![exesize3_26](image/exe3_26.png)

<br></br>

### <a name="3_27"></a> 27. Для каждой буквы алфавита из числа найденных величин, средних и балльных. (Т.е. среди всех студентов, чьё имя начинается на А (Алексей, Алина, Артур, Анджела) найти то, что указано в задании. Вывести на экран тех, средний балл проходит больше 3,6

#### `Запрос`

```SQL
SELECT LEFT(st.name::VARCHAR,1) kours, MIN(st.score), MAX(st.score), ROUND(AVG(st.score),2) sred
FROM students st
GROUP BY LEFT(st.name::VARCHAR,1)
HAVING MAX(st.score) > 3.6
```

#### `Вывод`

![exesize3_27](image/exe3_27.png)

<br></br>

### <a name="3_28"></a> 28. Для каждой фамилии характерна высокая степень тяжести и снижение балла. (Например, в университете учатся 4 Иванова (1-2-3-4). 1-2-3 учатся на 2 курсах и имеют средний балл 4.1, 4, 3.8, соответственно, а 4 Иванов учится на 3 курсах и имеет балл 4.5. На экране должно быть следующее:
курс | фамилия | max | min
-----|---------|-----|-----
2    | Иванов  | 4,1 | 3,8
3    | Иванов  | 4,5 | 4,5

#### `Запрос`

```SQL
SELECT 
  LEFT(st.n_group::VARCHAR,1) kours, 
  st.surname,
  MIN(st.score),
  MAX(st.score)
FROM students st
GROUP BY
  LEFT(st.n_group::VARCHAR,1),
  st.surname
```

#### `Вывод`

![exesize3_28](image/exe3_28.png)

<br></br>

### <a name="3_29"></a> 29. Для каждого года рождения подсчитывается количество хобби, наблюдаемое или занимавшееся студентами.

#### `Запрос`

```SQL
SELECT EXTRACT(YEAR FROM st.date_birth) birth, COUNT(*) col_hobby
FROM students st
INNER JOIN students_hobbies st_hb
ON st.id = st_hb.student_id
GROUP BY EXTRACT(YEAR FROM st.date_birth)
```

#### `Вывод`

![exesize3_29](image/exe3_29.png)

<br></br>

### <a name="3_30"></a> 30. Для каждой буквы алфавита в большинстве случаев возникает риск, связанный с хобби.

#### `Запрос`

```SQL
--непонял задание
```

#### `Вывод`

<br></br>

### <a name="3_31"></a> 31. Для каждого месяца из-за даты рождения студентов получаются средние баллы, которые являются хобби под названием «Футбол».

#### `Запрос`

```SQL
SELECT EXTRACT(MONTH FROM st.date_birth) month_, COUNT(st.id) col_st
FROM students st
INNER JOIN students_hobbies st_hb
ON st_hb.id = st.id
INNER JOIN hobby hb
ON st_hb.hobby_id = hb.id
GROUP BY EXTRACT(MONTH FROM st.date_birth), hb.name
HAVING hb.name = 'футбол'
```

#### `Вывод`

![exesize3_31](image/exe3_31.png)

<br></br>

### <a name="3_32"></a> 32. Вывести информацию о студентах, которые занимались или занимались хотя бы 1 хобби в следующем формате: Имя: Иван, фамилия: Иванов, группа: 1234
#### `Запрос`

```SQL
SELECT st.name, st.surname, st.n_group
FROM students st
WHERE
  st.id IN 
    (SELECT st.id
    FROM students st
    INNER JOIN students_hobbies st_hb
    ON st_hb.student_id = st.id
    INNER JOIN hobby hb
    ON st_hb.hobby_id = hb.id
    GROUP BY st.id)
```

#### `Вывод`

![exesize3_32](image/exe3_32.png)

<br></br>

### <a name="3_33"></a> 33. Внешний вид в какой-то по счёту символа встречается «ов». Если 0 (т.е. не встречается, то выведите на экран «не найдено»).
#### `Запрос`

```SQL
SELECT 
  st.surname,
  CASE
    WHEN POSITION('ов' IN st.surname) = 0
    THEN 'Не найдено'
  ELSE POSITION('ов' IN st.surname)::VARCHAR
  END
FROM students st
```

#### `Вывод`

![exesize3_33](image/exe3_33.png)

<br></br>

### <a name="3_34"></a> 34. Дополнить фамилию прямым символом # до 10 символов.

#### `Запрос`

```SQL
SELECT OVERLAY('##########' placing st.surname FROM 1)
FROM students st
```

#### `Вывод`

![exesize3_34](image/exe3_34.png)

<br></br>

### <a name="3_35"></a> 35. При помощи функции удалить все символы # из полученного запроса.

#### `Запрос`

```SQL
SELECT TRIM(TRAILING '#' FROM OVERLAY('##########' placing st.surname FROM 1)) surname
FROM students st
```

#### `Вывод`

![exesize3_35](image/exe3_35.png)

<br></br>

### <a name="3_36"></a> 36. Выведите на экран количество дней в прошлом 2018 году.

#### `Запрос`

```SQL
SELECT EXTRACT(DAY FROM '2019-01-01'::TIMESTAMP-'2018-01-01'::TIMESTAMP)
```

#### `Вывод`

![exesize3_36](image/exe3_36.png)

<br></br>

### <a name="3_37"></a> 37. Вывести на экран какого числа будет ближайшая суббота.

#### `Запрос`

```SQL
SELECT NOW()::DATE + (6-EXTRACT(DOW FROM NOW()))::INT суббота
```

#### `Вывод`

![exesize3_37](image/exe3_37.png)

<br></br>

### <a name="3_38"></a> 38. Выведите на экран век, а также какая сейчас неделя года и день года.

#### `Запрос`

```SQL
SELECT 
  EXTRACT(CENTURY FROM NOW()) cent, 
  EXTRACT(WEEK FROM NOW()) week,
  EXTRACT(DOY FROM NOW()) days
```

#### `Вывод`

![exesize3_38](image/exe3_38.png)

<br></br>

### <a name="3_39"></a> 39. Вы вводите всех студентов, которые занимались хотя бы 1 хобби. Выведите на экран Имя, Фамилию, Название хобби, а также надпись «занимается», если студент продолжает заниматься хобби в данный момент или «закончил», если уже не занимается.

#### `Запрос`

```SQL
SELECT 
  st.name, 
  st.surname,
  hb.name,
  CASE
    WHEN (st_hb.date_finish IS NULL) THEN 'Занимается'
    WHEN (st_hb.date_finish is NOT NULL) THEN 'Закончил'
  END status
FROM students st
INNER JOIN students_hobbies st_hb
ON st_hb.id = st.id
INNER JOIN hobby hb
ON st_hb.hobby_id = hb.id
```

#### `Вывод`

![exesize3_39](image/exe3_39.png)

<br></br>

### <a name="3_40"></a> 40. Для каждой группы приводят количество студентов учится на 5,4,3,2. Использовать обычное математическое округление. Итоговый результат должен выглядеть примерно так:


СЧЕТ  | 2222 | 3011 | 4011 | 4032 
----- | ---- | ---- | ---- | ---- 
2     | 0    | 0    | 0    | 1 
3     | 1    | 2    | 1    | 1 
4     | 4    | 3    | 3    | 3 
5     | 1    | 1    | 1    | 0 

#### `Запрос`

```SQL
SELECT 
  st.n_group, 
  COUNT(st.score) FILTER (WHERE ROUND(st.score) = 5) five,
  COUNT(st.score) FILTER (WHERE ROUND(st.score) = 4) four,
  COUNT(st.score) FILTER (WHERE ROUND(st.score) = 3) three,
  COUNT(st.score) FILTER (WHERE ROUND(st.score) = 2) two
FROM students st
GROUP BY
  st.n_group
```

#### `Вывод`

![exesize3_40](image/exe3_40.png)

<br></br>