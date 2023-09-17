$(document).ready(function() {
     getDocumentos()
    getExportBase64()
    $('#documentos').DataTable();
});
var b64Documents = ""
async function getExportBase64(){
    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    const request = await fetch("http://localhost:5000/document/export/obra/" + obraid, {
        method: 'GET',
    })
    console.log(request.status)
    const body = await request.json();
    if (request.status == 200){
        // Obtener una referencia al botón por su ID
        b64Documents = body.document
        var boton = document.getElementById("botonDescargar");

    // Habilitar el botón
        boton.removeAttribute("disabled");
    }


}
async function getDocumentos() {

    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    const request = await fetch("http://localhost:5000/document/" + obraid, {
        method: 'GET',
    })

    const documentos = await request.json();
    console.log(documentos)
    if (request.ok) {
        let listadoHtml = '';
            for (let documento of documentos) {

                let estado = ''
                if (documento.status == "DONE"){
                    estado += '<td bgcolor="#adebad">Revisado</td>';
                }
                if (documento.status == "WIP"){
                    estado += '<td bgcolor="#ffff99">Falta Revisar</td>';
                }
                if (documento.status == "TODO"){
                    estado += '<td bgcolor="#ffb3b3">Sin Revisar</td>';
                }

                let botonChecks = '<a onclick= "redirectRevisarChecks('+ documento.id +')" class="btn btn-success btn-icon-split">Revisar Checks</a>'
                let botonEliminar = '<a onclick = "eliminarDocumento('+documento.id+')"  class="btn btn-success btn-icon-split">Eliminar</a>'
              let documentoHtml = '<tr><td>' + documento.formulario.nombre + '</td><td>' + documento.formulario.informacion + '</td><td>' + documento.piso.numero + '</td>' + estado + '<td>' + botonChecks + botonEliminar +'</td></tr>';
              listadoHtml += documentoHtml;
              }
            
            
            document.querySelector('#documentosTBody').outerHTML = listadoHtml;
            
            }
    }


function redirectCrearDocumento(){

    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    window.location.href = `altaDePlan.html?obra_id=${obraid}`;
}

function redirectRevisarChecks(id){
    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    window.location.href = `revisarChecks.html?documento_id=${id}&obra_id=${obraid}`;
}

// Función para descargar el PDF en base64 al hacer clic en el botón
document.getElementById('botonDescargar').addEventListener('click', function() {
    // Supongamos que tienes el PDF codificado en base64 en una variable llamada 'pdfBase64'

    // Crea un enlace <a> temporal en el DOM
    var enlaceDescarga = document.createElement('a');

    // Establece el atributo href del enlace con el contenido base64
    enlaceDescarga.href = "data:application/pdf;base64," + b64Documents;

    // Establece el atributo download para que el navegador inicie la descarga en lugar de navegar a la URL
    enlaceDescarga.download = "exported_documents.pdf"; // Puedes cambiar el nombre del archivo si lo deseas

    // Oculta el enlace para que no sea visible en la página
    enlaceDescarga.style.display = 'none';

    // Agrega el enlace al DOM
    document.body.appendChild(enlaceDescarga);

    // Simula un clic en el enlace para iniciar la descarga
    enlaceDescarga.click();

    // Elimina el enlace del DOM después de la descarga
    document.body.removeChild(enlaceDescarga);
});

async function eliminarDocumento(id){

    const request = await fetch("http://localhost:5000/document/eliminar/"+ id, {
               method: 'DELETE',
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.ok) {
             alert("Documento Eliminado");
         }else{
             alert("Error eliminando el documento");
         }

         let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');
    window.location.href = `documentosDeObra.html?obra_id=${obraid}`;
  }