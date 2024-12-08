SELECT c.name AS child_name
FROM children c
JOIN gifts g
ON c.child_id = g.child_id
WHERE g.price > (SELECT AVG(price) FROM gifts)
ORDER BY g.price
LIMIT 1;
