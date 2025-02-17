SELECT snippets.*, logging.*
FROM snippets
INNER JOIN logging ON snippets.log_id = logging.logging_id
WHERE snippets.log_id IS NOT NULL;