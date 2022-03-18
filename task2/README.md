# Задание 2

Однотабличные запросы

1. [Вывести всеми возможными способами имена и фамилии студентов, средний балл которых от 4 до 4.5](#1)
2. [Познакомиться с функцией CAST. Вывести при помощи неё студентов заданного курса (использовать Like)](#2)
3. [Вывести всех студентов, отсортировать по убыванию номера группы и имени от а до я](#3)
4. [Вывести студентов, средний балл которых больше 4 и отсортировать по баллу от большего к меньшему](#4)
5. [Вывести на экран название и риск футбола и хоккея](#5)
6. [Вывести id хобби и id студента которые начали заниматься хобби между двумя заданными датами (выбрать самим) и студенты должны до сих пор заниматься хобби](#6)


##  <a name="1"></a> Вывести всеми возможными способами имена и фамилии студентов, средний балл которых от 4 до 4.5

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where st.score <=4.5 and st.score >= 4
```

### Вывод:

![exesize1](image/exesize1.png)

## <a name="2"></a> Познакомиться с функцией CAST. Вывести при помощи неё студентов заданного курса (использовать Like)

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

## <a name="3"></a> Вывести всех студентов, отсортировать по убыванию номера группы и имени от а до я

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
order by st.n_group desc, st.name 
```

### Вывод

![exesize3](image/exesize3.png)

## <a name="4"></a> Вывести студентов, средний балл которых больше 4 и отсортировать по баллу от большего к меньшему

### Запрос

```SQL
select st.name, st.surname, st.n_group, st.score
from students st
where score > 4
order by st.score desc
```

### Вывод

![exesize4](image/exesize4.png)

## <a name="5"></a> Вывести на экран название и риск футбола и хоккея

### Запрос

```SQL
select hb.name, hb.risk
from hobby hb
where name = 'футбол' or name = 'хоккей'
```

### Вывод

![exesize4](image/exesize5.png)

## <a name="6"></a> Вывести id хобби и id студента которые начали заниматься хобби между двумя заданными датами (выбрать самим) и студенты должны до сих пор заниматься хобби

### Запрос

```SQL