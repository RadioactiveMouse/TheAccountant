MyMoney
========

Aims
------

To produce a daemon that watches a directory for new files from Barclays export data function and convert them into a more human readable and searchable  form. 

Must be able to :

Read Barclays exported data
Process it based on a config file
Automatically check for new files in the directory at specified intervals
Access the interface from localhost:port
Graph the data contained in a meaningful way


Method
-------

Abstract a graph library using Go and D3.js in order to ensure modularity.


