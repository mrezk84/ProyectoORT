$(document).ready(function() {
    getChecks()

    $('#checksDeDocumento').DataTable();
});




async function getChecks() {

    let documentoid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    documentoid = searchParams.get('documento_id');


    const request = await fetch("http://localhost:5000/checks/document/" + documentoid, {
        method: 'GET',
    })

    const checks = await request.json();
    console.log(checks)
    if (request.ok) {
        let listadoHtml = '';
            for (let check of checks) {

              let input = '<input type="text" class="form-control form-control-user" id="txt' + check.id + '" placeholder="'+ check.observaciones +'"></input>'
              let select = '<select id="select'+ check.id +'"><option value="NO CONFORME">NO CONFORME</option><option value="CONFORME">CONFORME</option></select>'

              let controlHtml = '<tr><td>'+ check.id +'</td><td>' + check.estado + '</td><td>' + check.observaciones + '</td><td>' + select 
              + '</td><td>'+ input +'</td></tr>';
              listadoHtml += controlHtml;
              }
            
            
            document.querySelector('#controlesTBody').outerHTML = listadoHtml;
            
            }
    }


function redirectCrearControl(){
    window.location.href = `altaControl.html`;
}