-- Mostrar el título y el nombre del género de todas las series.
-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.



SELECT s.title , g.name  FROM series s 
INNER JOIN genres g on s.genre_id = g.id


SELECT e2.title episodio, a.first_name nombre, a.last_name apellido  
FROM actors a 
INNER JOIN actor_episode ae ON  ae.actor_id = a.id
INNER JOIN episodes e2 ON e2.id = ae.episode_id 
ORDER BY e2.title 

SELECT s.title , COUNT(*) as totalDeTemporadas
FROM series s 
INNER JOIN seasons s2 ON s2.serie_id = s.id 
group by s.title 


SELECT g.name, COUNT(*) as totalPeliculas  
FROM genres g 
INNER JOIN movies m ON m.genre_id = g.id
group by g.id HAVING totalPeliculas > 3


select DISTINCT (a.first_name), a.last_name 
FROM actors a 
INNER JOIN actor_movie am ON a.id = am.actor_id 
INNER JOIN movies m ON m.id = am.movie_id 
WHERE m.title like "% guerra de las galaxias%"