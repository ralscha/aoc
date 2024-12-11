SELECT date
FROM Drinks
GROUP BY date
HAVING SUM(CASE WHEN drink_name = 'Hot Cocoa' THEN quantity ELSE 0 END) = 38
   AND SUM(CASE WHEN drink_name = 'Peppermint Schnapps' THEN quantity ELSE 0 END) = 298
   AND SUM(CASE WHEN drink_name = 'Eggnog' THEN quantity ELSE 0 END) = 198;