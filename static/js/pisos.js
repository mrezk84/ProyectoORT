$(document).ready(function() {
    getPisosByObra()
});



let obra_id = null;
async function getPisosByObra() {
    const url = new URL(document.URL);
    const searchParams = url.searchParams;

    obra_id = searchParams.get('obra_id');

    if (obra_id != null) {
        console.log(obra_id)
    }
    const request = await fetch(`obra/${obra_id}`, {
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
                    <th>${piso.nroPiso}</th>
                    <th><button onclick="redirectFormularios(${obra_id},${piso.id})">Ver Formularios</button></th>
                </tr>
                `
            });
            document.getElementById("pisosTBody").innerHTML = text;

    }
}

async function agregarPisoObra(){
    let piso = document.getElementById("idPisoAgregar").value;

    console.log({
        "nroPiso": piso,
    })
    const request = await fetch(`piso/${obra_id}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "nroPiso": piso,
        })
    });
    if (request.status == 200){
        alert("Piso agregado correctamente");
        window.location.reload();
    }
    else{
        alert("Error al agregar el piso");
    }
}

function redirectFormularios(obraID,pisoID) {
    window.location.href = `formulariosDeObra.html?obra_id=${obraID}&piso_id=${pisoID}`;
}