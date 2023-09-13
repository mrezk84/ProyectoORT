$(document).ready(function() {
    getPisos()
});



async function getPisos() {
    const request = await fetch("http://localhost:5000/pisos", {
        method: 'GET',
    })
    let pisos = await request.json()
    console.log(pisos)
    if (request.status == 200) {
        let listadoHtml = '';
        for (let piso of pisos) {

            let botonUpdate = '<a onclick = "redirectUpdate('+piso.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Update</a>';
            let botonDelete = '<a onclick = "eliminarPiso('+piso.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Delete</a>';
          let pisoHtml = '<td>' + piso.id + '</td><td>' + piso.numero + '</td><td>' + botonUpdate + botonDelete+ '</td></tr>';
          listadoHtml += pisoHtml;
          }
            document.getElementById("pisosTBody").outerHTML = listadoHtml;

    }
}

async function redirectAltaPiso() {
    window.location.href = `altaPiso.html`;
}

async function redirectUpdate(id) {
    window.location.href = `pisoUpdate.html?piso_id=${id}`;
}

async function eliminarPiso(id){

    const request = await fetch("http://localhost:5000/controles/"+ id +"/", {
               method: 'DELETE',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.ok) {
             alert("Piso Eliminado");
         }else{
             alert("Error eliminando el piso");
         }
  }

// async function agregarPisoObra(){
//     let piso = document.getElementById("idPisoAgregar").value;

//     console.log({
//         "nroPiso": piso,
//     })
//     const request = await fetch(`piso/${obra_id}`, {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json'
//         },
//         body: JSON.stringify({
//             "nroPiso": piso,
//         })
//     });
//     if (request.status == 200){
//         alert("Piso agregado correctamente");
//         window.location.reload();
//     }
//     else{
//         alert("Error al agregar el piso");
//     }
// }

// function redirectFormularios(obraID,pisoID) {
//     window.location.href = `formulariosDeObra.html?obra_id=${obraID}&piso_id=${pisoID}`;
// }