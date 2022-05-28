# Сервис складского учёта

1. Администратор. Добавление нового оператора, пользователя
(работник склада)
2. Оператор. Добавление нового поставщика в систему
3. Оператор. Добавление нового товара в систему (название, вес,
количество, объём, от кого) (Поставка товара)
4. Оператор. Добавление информации по отгрузки товара (кому, что,
сколько, когда)
5. Пользователь. Поиск товара
6. Пользователь. Просмотр задач на сегодня (с поиском по различным
параметрам, пагинацией)
7. Оператор. Назначить задачу на пользователя
8. Пользователь. Взять задачу в работу
9. Администратор. Просмотр выполненных заказов (можно задать дату,
id, поиск по названию поставщика, поиск по названию адресата доставка)
10. Администратор. Блокировка оператора, пользователя.

---------------------------------------------------------------------

## Таблицы

### Товары

```SQL
CREATE TABLE  products(
	-- название, вес,количество, объём, от кого
	id serial PRIMARY KEY,
	name varchar(255) not null,
	weight_kg decimal not null,
	amount int not null,
	size decimal not null,
	from_whom varchar(20) not null
)

CREATE sequence products_seq;
```

### Поставщики

```SQL
CREATE TABLE  suppliers(
	id serial PRIMARY KEY,
	name varchar(255) not null
);

CREATE sequence suppliers_seq;
```

### Транзит

```SQL
CREATE TABLE  transit(
	id serial PRIMARY KEY,
	id_product int not null,
	id_provider int not null,
	amount decimal not null,
    adress varchar(255) not null, 
	departure_date date not null,
    status int not null -- 0 доступна, 1 в работе, 2 выполнена
);

CREATE sequence transit_seq;

alter table transit add
constraint status_transit_CHECK
check (status >= 0 and status <= 2);

ALTER TABLE transit ADD CONSTRAINT product_transit_FK
FOREIGN KEY (id_product)
REFERENCES products (id);

ALTER TABLE transit ADD CONSTRAINT provider_transit_FK
FOREIGN KEY (id_provider)
REFERENCES suppliers (id);
```

### Пользователи

```SQL
CREATE TABLE  users(
	id serial PRIMARY KEY,
	name varchar(255) not null,
	surname varchar(255) not null,
	post varchar(255) not null,
	access int not null -- 0 администратор, 1 оператор, 2 пользователь --
);

CREATE sequence users_seq;

alter table users add
constraint access_CHECK
check (access >= 0 and access <= 2);
```

### Задачи

```SQL
CREATE TABLE  tasks(
	id serial PRIMARY KEY,
	name varchar(255) not null,
	price decimal not null,
	status int not null -- 0 доступна, 1 взята, 2 выполнена
);

CREATE sequence tasks_seq;

alter table tasks add
constraint status_CHECK
check (status >= 0 and status <= 2);
```

### Работа

```SQL
CREATE TABLE  work(
	id serial PRIMARY KEY,
	id_task int not null,
	id_user int not null
);

CREATE sequence work_seq;

ALTER TABLE work ADD CONSTRAINT work_task_FK
FOREIGN KEY (id_task)
REFERENCES tasks (id);

ALTER TABLE work ADD CONSTRAINT work_user_FK
FOREIGN KEY (id_user)
REFERENCES users (id);
```

## Администратор

1. Администратор. Добавление нового оператора, пользователя
(работник склада)
2. Администратор. Просмотр выполненных заказов (можно задать дату,
id, поиск по названию поставщика, поиск по названию адресата доставка)
3. Администратор. Блокировка оператора, пользователя.

```SQL
INSERT INTO users (name, surname, post, access)
VALUES (?, ?, ?, ?); -- или VALUES (?) для добавления массивом
```

```SQL
select sup.name, pr.name, tr.amount, tr.adress, tr.departure_date
from transit tr
INNER JOIN products pr
ON pr.id = tr.id_product
INNER JOIN suppliers sup
ON sup.id = tr.id_provider
where  status=2 and (pr.name = ? or tr.id = ? or tr.adress = ? or tr.departure_date = ?) -- нужное(ненужное) добавить(убрать) через php легко делается
```

```SQL
delete from users
where name=? and surname=? -- или id = ?
```

## оператор

1. Оператор. Добавление нового поставщика в систему
2. Оператор. Добавление нового товара в систему (название, вес,
количество, объём, от кого) (Поставка товара)
3. Оператор. Добавление информации по отгрузки товара (кому, что,
сколько, когда)
4. Оператор. Назначить задачу на пользователя

```SQL
INSERT INTO users (name, surname, post, access)
VALUES (?, ?, ?, 2); -- или VALUES (?) для добавления массивом
```

```SQL
insert into products (name, weight_kg, amount, size, from_whom)
value (?, ?, ?, ?, ?) -- или VALUES (?) для добавления массивом
```

```SQL
insert into transit (id_product, id_provider, amount, adress, departure_date)
value (?, ?, ?, ?, ?) -- или VALUES (?) для добавления массивом
```

```SQL
insert into work (id_task, id_user)
value(?, ?)
```

## Пользователь 

1. Пользователь. Поиск товара
2. Пользователь. Просмотр задач на сегодня (с поиском по различным параметрам, пагинацией)
3. Пользователь. Взять задачу в работу

```SQL
select name, weight_kg, amount, size, from_whom
from products
where name=? or amount=? -- нужное(ненужное) добавить(убрать) через php легко делается
```

```SQL
select name, price, status
from tasks
where status != 2
```

```SQL
insert into work (id_task, id_user)
value(? , $_SESSION[id]) -- второй параметр определяется сессией и ставится автоматически (php)
```