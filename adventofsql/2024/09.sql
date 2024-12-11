SELECT 
    reindeer_name, 
    ROUND(MAX(avg_speed), 2) AS top_speed
FROM (
    SELECT 
        r.reindeer_name,
        ts.reindeer_id,
        ts.exercise_name,
        AVG(ts.speed_record) AS avg_speed
    FROM 
        Training_Sessions ts
    JOIN 
        Reindeers r ON ts.reindeer_id = r.reindeer_id
    WHERE 
        r.reindeer_name != 'Rudolf'  -- Excluding Rudolf
    GROUP BY 
        r.reindeer_name, ts.reindeer_id, ts.exercise_name
) AS AverageSpeeds
GROUP BY 
    reindeer_name
ORDER BY 
    top_speed DESC
LIMIT 3;
