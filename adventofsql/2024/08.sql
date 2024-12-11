WITH RECURSIVE hierarchy_cte AS (
    SELECT 
        staff_id,
        staff_name,
        1 AS level,
        CAST(staff_id AS VARCHAR) AS path
    FROM 
        staff
    WHERE 
        manager_id IS NULL
    
    UNION ALL
    
    SELECT 
        s.staff_id,
        s.staff_name,
        h.level + 1 AS level,
        CONCAT(h.path, ',', s.staff_id) AS path
    FROM 
        staff s
    INNER JOIN 
        hierarchy_cte h ON s.manager_id = h.staff_id
)
SELECT 
    max(level)
FROM 
    hierarchy_cte
