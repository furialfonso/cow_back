insert into c_group (code,debt)values('you&I',0);
insert into c_user(name,second_name,last_name,second_last_name,email,nick_name)values('Andres','Felipe','Alfonso','Ortiz','afelipealfonso@gmail.com','furialfonso');
insert into c_user(name,second_name,last_name,second_last_name,email,nick_name)values('Johanna','Lorena','Marquez','Torres','jolomato@gmail.com','joha');
insert into c_team(group_id,user_id)values(1,1);
insert into c_team(group_id,user_id)values(1,2);
insert into c_pay(team_id,description,value)values(1,'hamburgesas',20200);
insert into c_pay(team_id,description,value)values(2,'qbano',55000);
insert into c_pay(team_id,description,value)values(2,'corral',80000);
insert into c_pay(team_id,description,value)values(2,'a kikiri ki',30000);