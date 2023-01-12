# DBGO

## Proposal
* code generator based struct to access database easily.


~~~mermaid
flowchart LR

goDriver[/Go stdlib databse/sql/]
goStruct[Go struct]
mysql[(MySQL)]

subgraph CLI
    Config
    Model
    CodeGenerator
    
    Config --> CodeGenerator
    Model --> CodeGenerator
end    

subgraph WriteCode
    goStruct
end


subgraph Runtime
    SQL
    Values
    
    subgraph Executor
        goDriver
    end
end

CodeGenerator --> goStruct

goStruct --> SQL
goStruct --> Values

SQL --> goDriver
Values --> goDriver

goDriver <---> mysql 



~~~