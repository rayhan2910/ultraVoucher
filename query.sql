SELECT 
s.id_name,
s.name,
a.name as parent_name
from testt s
LEFT JOIN 
testt a ON s.id_parent = a.id_name

