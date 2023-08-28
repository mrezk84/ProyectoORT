$(document).ready(function() {
    getUsuarios()
});




async function getUsuarios() {
    const request = await fetch("http://localhost:8080/usuarios", {
        method: 'GET',
    })

    const usuarios = await request.json();
    console.log(usuarios)
    if (request.ok) {
        let listadoHtml = '';
            for (let usuario of usuarios) {
                listadoHtml +=
                `
                 <option value="${usuario.ID}">${usuario.Name}</option>
                `
        };

        document.getElementById("select").innerHTML = listadoHtml;

    }
}