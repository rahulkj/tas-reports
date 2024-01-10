scripts = [
    "https://fastly.jsdelivr.net/npm/jquery@3.7.1/dist/jquery.min.js",
    "https://fastly.jsdelivr.net/npm/echarts@5.4.3/dist/echarts.min.js",
    "https://code.jquery.com/jquery-3.7.0.js",
    "https://cdn.datatables.net/1.13.7/js/jquery.dataTables.js",
    "https://cdn.datatables.net/buttons/2.4.2/js/dataTables.buttons.min.js",
    "https://cdnjs.cloudflare.com/ajax/libs/jszip/3.10.1/jszip.min.js", // For exporting in excel
    "https://cdn.datatables.net/buttons/2.4.2/js/buttons.html5.min.js", // For exporting in csv and copy operations
    "https://cdn.datatables.net/colreorder/1.7.0/js/dataTables.colReorder.min.js",
    "https://cdn.datatables.net/buttons/2.4.2/js/buttons.colVis.min.js",
    "https://cdn.datatables.net/select/1.7.0/js/dataTables.select.min.js",
]

scripts.forEach(script => {
    var node = document.createElement('script');
    node.type = "text/javascript";
    node.src = script;
    document.body.appendChild(node);    
});
