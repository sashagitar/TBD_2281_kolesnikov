# Задание 2

Однотабличные запросы

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

##  <a name="1_0"></a> Вывести всеми возможными способами имена и фамилии студентов, средний балл которых от 4 до 4.5

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where st.score <=4.5 and st.score >= 4
```

### Вывод:

![exesize1](image/exesize1.png)

## <a name="2_0"></a> Познакомиться с функцией CAST. Вывести при помощи неё студентов заданного курса (использовать Like)

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where cast (st.n_group as varchar) like '2281'

-- или

select st.name, st.surname, st.n_group, st.score
from students st
where st.n_group::varchar like '2281'
```

### Вывод

![exesize2](image/exesize2.png)

## <a name="3_0"></a> Вывести всех студентов, отсортировать по убыванию номера группы и имени от а до я

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
order by st.n_group desc, st.name 
```

### Вывод

![exesize3](image/exesize3.png)

## <a name="4_0"></a> Вывести студентов, средний балл которых больше 4 и отсортировать по баллу от большего к меньшему

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4
order by st.score desc
```

### Вывод

![exesize4](image/exesize4.png)

## <a name="5_0"></a> Вывести на экран название и риск футбола и хоккея

### Запрос

```SQL
select hb.name, hb.risk
from hobby hb
where name = 'футбол' or name = 'хоккей'
```

### Вывод

![exesize5](image/exesize5.png)

## <a name="6_0"></a> Вывести id хобби и id студента которые начали заниматься хобби между двумя заданными датами (выбрать самим) и студенты должны до сих пор заниматься хобби

### Запрос

```SQL
select st.student_id, st.hobby_id, st.date_start, st.date_finish
from students_hobbies st
where  date_start >= '2002-04-10' and date_finish is null;
```

### Вывод

![exesize6](image/exesize6.png)

## <a name="7_0"></a> Вывести студентов, средний балл которых больше 4.5 и отсортировать по баллу от большего к меньшему

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4.5
order by st.score 
```

### Вывод

![exesize7](image/exesize7.png)

## <a name="8_0"></a> Из запроса №7 вывести несколькими способами на экран только 5 студентов с максимальным баллом

### Запрос

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

### Вывод

![exesize8](image/exesize8_1.png)

![exesize8.2](image/exesize8_2.png)

![exesize8.3](image/exesize8_3.png)

![exesize8.4](image/exesize8_4.png)

## <a name="9_0"></a> Выведите хобби и с использованием условного оператора сделайте риск словами:

### Запрос

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

### Вывод

![exesize8.4](image/exesize9.png)

## <a name="10_0"></a> Вывести 3 хобби с максимальным риском

### Запрос

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

### Вывод

![exesize8.4](image/exesize10.png)
