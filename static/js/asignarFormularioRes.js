$(document).ready(function() {
    getUsuarios()
});




async function getUsuarios() {
    const request = await fetch("http://localhost:5000/usuarios", {
        method: 'GET',
    })

    const usuarios = await request.json();
    console.log(usuarios)
    if (request.ok) {
        let listadoHtml = '';
            for (let usuario of usuarios) {
                listadoHtml +=
                `
                 <option value="${usuario.id}">${usuario.name}</option>
                `
        };

        document.getElementById("select").innerHTML = listadoHtml;

    }
}