window.addEventListener('DOMContentLoaded',()=>{loadUsers()})
function abrirModal() {
  document.getElementById("miModal").style.display = "block";
  const obtenerPokemons = async()=>{
    try{
        const res = await fetch('http://localhost:8080/')
        mode: 'no-cors'
        const data = await res.json()
        console.log(data.results)
        const Nomnbres = data.results.map( poke => poke.name)
        console.log(Nomnbres)
    }catch(error){
        console.log(error)
    }
}
obtenerPokemons()
  /*fetch('http://localhost:8080/')

  .then(response => response.json())
  .then(data => {
    // Acceder a los datos JSON y mostrarlos en el frontend
    const container = document.getElementById('datos-container');
    
    // Limpiar el contenido existente
    container.innerHTML = '';
    
    // Recorrer los datos y crear elementos HTML para mostrarlos
    data.forEach(item => {
      const elemento = document.createElement('p');
      elemento.textContent = `${item.Title}: ${item.Numero}`;
      container.appendChild(elemento);
    });
  })
  .catch(error => {
    console.error('Error:', error);
  });*/
  }
  
  function cerrarModal() {
    document.getElementById("miModal").style.display = "none";
  }
/*-------------- menu desplegable-------------------- */
function mostrar(){
  document.getElementById("opciones").style.display = "block";
 }

 function ocultar(){
  document.getElementById("opciones").style.display = "none";
 }
 
  /*hacer que ak dar click en entrada se despliegue el menu en la pagina eliminar*/
  const buttom = document.querySelector('.button')
  const nav = document.querySelector('.nav')

  buttom.addEventListener('click', ()=>{
      nav.classList.toggle('activo')
  }
  )
 