function abrirModal() {
    document.getElementById("miModal").style.display = "block";
  }
  
  function cerrarModal() {
    document.getElementById("miModal").style.display = "none";
  }
  /*hacer que ak dar click en entrada se despliegue el menu*/
  const buttom = document.querySelector('.button')
  const nav = document.querySelector('.nav')

  buttom.addEventListener('click', ()=>{
      nav.classList.toggle('activo')
  }
  )