$(document).ready(function() {
    getPisos()
});



async function getPisos() {
    const request = await fetch("http://localhost:5000/pisos", {
        method: 'GET',
    })
    let response = await request.json()
    console.log(response)
    if (request.status == 200) {
            let text = ``;
            response.pisosObra.forEach(piso => {
            text +=
            `
                <tr>
                    <th>${piso.id}</th>
                    <th>${piso.numero}</th>
                    <th>${piso.numero}</th>
                </tr>
                `
            });
            document.getElementById("pisosTBody").innerHTML = text;

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