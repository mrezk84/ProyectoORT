$(document).ready(function() {
    getObrasByUser()
});




async function getObrasByUser() {
    const request = await fetch("http://localhost:5000/obras", {
        method: 'GET',
    })

    const obras = await request.json();
    console.log(obras)
    if (request.ok) {
        let listadoHtml = '';
            for (let obra of obras) {
                listadoHtml +=
                `
                 <tr>
                    <th>${obra.ID}</th>s
                    <th>${obra.Nombre}</th>
                    <th><button onclick="redirectPisos(${obra.ID})">Ver seguimiento</button></th>
                    <th><button onclick="redirectDocumentos(${obra.ID})">Ver documentos</button></th>
                </tr>
                `
        };

        document.querySelector('#obrasTBody').outerHTML = listadoHtml;

    }
}

function redirectPisos(id) {
    window.location.href = `altaSeguimiento.html?obra_id=${id}`;
}

function redirectDocumentos(id) {
    window.location.href = `documentosDeObra.html?obra_id=${id}`;
}

function redirectAltaObra() {
    window.location.href = `altaObra.html`;
}