function toggleDropdown() {
  var checkbox = document.getElementById('modulCheckbox');
  var dropdown = document.getElementById('modulDropdown');
  
  if (checkbox.checked) {
    dropdown.style.display = 'block';
  } else {
    dropdown.style.display = 'none';
  }
}

