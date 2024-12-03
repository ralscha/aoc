WITH cleaned_letters AS (
    SELECT id, value
    FROM letters_a
    WHERE value between ascii('a') AND ascii('z')
    OR value between ascii('A') AND ascii('Z')
    OR chr(value) in (' ', '!', ',', '.')

    UNION ALL

    SELECT id, value
    FROM letters_b
    WHERE value between ascii('a') AND ascii('z')
    OR value between ascii('A') AND ascii('Z')
    OR chr(value) in (' ', '!', ',', '.')
)
SELECT string_agg(CHR(value), '' ORDER BY id) AS decoded_message
FROM cleaned_letters;
