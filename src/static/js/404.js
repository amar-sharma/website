var ResponsiveVenDiag = function() {
  function init() {
    uiBindings();
  }

  function uiBindings() {
    window.addEventListener('resize', () => {
      let wrap = document.querySelector('.fof-wrapper');
      let baseTranslateVal = 'translateX(-50%) translateY(-50%)';
      let winWidth = this.innerWidth;

      if (winWidth < 1000) {
        let calcPerc = winWidth / 10;
        wrap.style.transform = baseTranslateVal + ' scale(.' + calcPerc + ')';
      } else {
        wrap.style.transform = baseTranslateVal + ' scale(1)';
      }
    });
  }

  return {
    init: init
  }
}();