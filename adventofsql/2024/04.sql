WITH tag_changes AS (
    SELECT toy_id,
           ARRAY(
               SELECT UNNEST(new_tags)
               EXCEPT
               SELECT UNNEST(previous_tags)
           ) AS added_tags,
           ARRAY(
               SELECT UNNEST(previous_tags)
               INTERSECT
               SELECT UNNEST(new_tags)
           ) AS unchanged_tags,
           ARRAY(
               SELECT UNNEST(previous_tags)
               EXCEPT
               SELECT UNNEST(new_tags)
           ) AS removed_tags
    FROM toy_production
)
SELECT toy_id,
       (SELECT COUNT(*) FROM UNNEST(added_tags)) AS added,
       (SELECT COUNT(*) FROM UNNEST(unchanged_tags)) AS unchanged,
       (SELECT COUNT(*) FROM UNNEST(removed_tags)) AS removed
FROM tag_changes
ORDER BY added DESC
LIMIT 1;
