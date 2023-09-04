# Gogensql.
O propósito deste programa é gerar um script SQL para atualização de tabalas a partir de um arquivo excel. 

# Como usar?
O programa é para ser executado via terminal de comando e passando os devidos argumentos abaixo:
 - Coluna da data: Digite, exatamente, a coluna que contém a data da meta, exemplo: S4.
 - Coluna do orçamento: Digite, apenas, a coluna da orçamento, exemplo: S.
 - Coluna da margem: Digite, apenas, a coluna da margem, exemplo: T.
 - Arquivo excel: Digite, entre aspas, o caminho do arquivo excel, exemplo: 'C:\Arquivo\Excel\Nome do arquivo.xlsx'

O comando do terminal deverá ser algo parecido com isto:
```
gogensql.exe S4 S T 'Nome do arquivo.xlsx'
```