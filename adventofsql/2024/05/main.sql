WITH production_with_previous AS (
    SELECT
        tp.production_date,
        tp.toys_produced,
        LAG(tp.toys_produced) OVER (ORDER BY tp.production_date) AS previous_day_production
    FROM
        toy_production tp
),
production_changes AS (
    SELECT
        production_date,
        toys_produced,
        previous_day_production,
        (toys_produced - previous_day_production) AS production_change,
        CASE 
            WHEN previous_day_production IS NOT NULL THEN
                ROUND(((toys_produced - previous_day_production) * 100.0) / previous_day_production, 2)
            ELSE NULL
        END AS production_change_percentage
    FROM
        production_with_previous
)
SELECT
    production_date,
    toys_produced,
    previous_day_production,
    production_change,
    production_change_percentage
FROM
    production_changes
where 
    production_change_percentage is not null 
ORDER BY
    ABS(production_change_percentage) DESC;
