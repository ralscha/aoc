WITH parsed_data AS (
    SELECT
        id,
        CASE
            WHEN menu_data::TEXT LIKE '%version="1.0"%' THEN
                CAST((xpath('//total_count/text()', menu_data)::TEXT[])[1] AS INT)
            WHEN menu_data::TEXT LIKE '%version="2.0"%' THEN
                CAST((xpath('//total_guests/text()', menu_data)::TEXT[])[1] AS INT)
            ELSE 0
        END AS guest_count,
        UNNEST(
            CASE
                WHEN menu_data::TEXT LIKE '%version="1.0"%' THEN
                    ARRAY(SELECT xpath('//food_item_id/text()', menu_data)::TEXT[])
                WHEN menu_data::TEXT LIKE '%version="2.0"%' THEN
                    ARRAY(SELECT xpath('//food_item_id/text()', menu_data)::TEXT[])
                ELSE ARRAY[]::TEXT[]
            END
        )::INT AS food_item_id
    FROM christmas_menus
),
filtered_data AS (
    SELECT food_item_id
    FROM parsed_data
    WHERE guest_count > 78
),
dish_frequency AS (
    SELECT
        food_item_id,
        COUNT(*) AS frequency
    FROM filtered_data
    GROUP BY food_item_id
)
SELECT food_item_id
FROM dish_frequency
ORDER BY frequency DESC
LIMIT 1;
