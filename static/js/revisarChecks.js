$(document).ready(function() {
    getChecks()

    $('#checksDeDocumento').DataTable();
});




async function getChecks() {

    let documentoid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    documentoid = searchParams.get('documento_id');


    const request = await fetch("http://34.192.187.56:5000/checks/document/" + documentoid, {
        method: 'GET',
    })

    const checks = await request.json();
    console.log(checks)
    if (request.ok) {
        let listadoHtml = '';
            for (let check of checks) {

              let input = '<input type="text" class="form-control form-control-user" id="txt' + check.id + '" placeholder="'+ check.observaciones +'"></input>';
              let select = '<select id="select'+ check.id +'"><option value="NO CONFORME">NO CONFORME</option><option value="CONFORME">CONFORME</option></select>';

              let checkHtml = '<tr><td>'+ check.id +'</td><td>' + check.estado + '</td><td>' + check.observaciones + '</td><td>' + select 
              + '</td><td>'+ input +'</td><td><button onclick="Actualizar('+ check.id +')" class="btn btn-primary btn-user">Actualizar</button></td></tr>';
              listadoHtml += checkHtml;
              }
            
            console.log(listadoHtml)
            document.querySelector('#checksTBody').outerHTML = listadoHtml;
            
            }
    }


async function Actualizar(id){
    let datos = {};
    datos.observaciones = document.getElementById('txt' + id).value;
    datos.estado = document.getElementById('select' + id).value;

    const request = await fetch("http://34.192.187.56:5000/checks/" + id, {
               method: 'PUT',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 201) {
             alert("check actualizado correctamente");
         }else{
             alert("Error actualizando el check");
         }
         let documentoid = null;
         const url = new URL(document.URL);
         const searchParams = url.searchParams;
         documentoid = searchParams.get('documento_id');
         let obraid = null;
         const url2 = new URL(document.URL);
         const searchParams2 = url2.searchParams;
         obraid = searchParams2.get('obra_id');
         window.location.href = `revisarChecks.html?documento_id=${documentoid}&obra_id=${obraid}`;
}

function Volver() {
    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    window.location.href = `documentosDeObra.html?obra_id=${obraid}`;
}