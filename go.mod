module <%= params.namespace ? `${params.namespace}/${projectName}` : projectName %>

go 1.22.3

require github.com/joho/godotenv v1.5.1 // indirect
