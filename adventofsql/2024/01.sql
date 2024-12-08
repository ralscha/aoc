SELECT
    c.name,
    (wishes->>'first_choice') AS primary_wish,
    (wishes->>'second_choice') AS backup_wish,
    (wishes->'colors'->0)::text AS favorite_color,
    json_array_length(wishes->'colors') AS color_count,
    CASE
        WHEN t.difficulty_to_make = 1 THEN 'Simple Gift'
        WHEN t.difficulty_to_make = 2 THEN 'Moderate Gift'
        ELSE 'Complex Gift'
    END AS gift_complexity,
    CASE t.category
        WHEN 'outdoor' THEN 'Outside Workshop'
        WHEN 'educational' THEN 'Learning Workshop'
        ELSE 'General Workshop'
    END AS workshop_assignment
FROM
    children c
JOIN
    wish_lists w ON c.child_id = w.child_id
JOIN
    toy_catalogue t ON t.toy_name = (wishes->>'first_choice')
ORDER BY
    c.name ASC
LIMIT 5;
