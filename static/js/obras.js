$(document).ready(function() {
    getObrasByUser()
});




async function getObrasByUser() {
    const request = await fetch('obra', {
        method: 'GET',
    })
    let response = await request.json()
    if (request.status == 200) {
        let text = ``;
        response.forEach(obra => {
            text +=
                `
                 <tr>
                    <th>${obra.id}</th>
                    <th>${obra.nombre}</th>
                    <th>${obra.pisosObra.length}</th>
                    <th><button onclick="redirectPisos(${obra.id})">Ver pisos</button></th>
                </tr>
                `
        });

        document.getElementById("obrasTBody").innerHTML = text;

    }
}

function redirectPisos(id) {
    window.location.href = `pisos.html?obra_id=${id}`;
}

function redirectAltaObra() {
    window.location.href = `altaObra.html`;
}