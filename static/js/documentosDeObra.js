$(document).ready(function() {
    // getDocumentos()

    $('#documentos').DataTable();
});

async function getControles() {

    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    const request = await fetch("http://localhost:5000/documentos", {
        method: 'GET',
    })

    const documentos = await request.json();
    console.log(documentos)
    if (request.ok) {
        let listadoHtml = '';
            for (let documento of documentos) {
  
              let documentoHtml = '<tr><td>'+ documento.ID +'</td><td>' + documento.Formulario + '</td><td>' + documento.Piso + '</td></tr>';
              listadoHtml += documentoHtml;
              }
            
            
            document.querySelector('#documentosTBody').outerHTML = listadoHtml;
            
            }
    }


function redirectCrearDocumento(){
    window.location.href = `altaDePlan.html`;
}