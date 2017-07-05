/*
 * ScaleAndExecution
 * https://github.com/code4nagoya/ScaleAndExecution
 *
 * Copyright (c) 2017 kouichi ume
 * Licensed under the Copyright license.
 */

(function($) {

  // Collection method.
  $.fn.ScaleAndExecution = function() {
    return this.each(function(i) {
      // Do something awesome to each selected element.
      $(this).html('awesome' + i);
    });
  };

  // Static method.
  $.ScaleAndExecution = function(options) {
    // Override default options with passed-in options.
    options = $.extend({}, $.ScaleAndExecution.options, options);
    // Return something awesome.
    return 'awesome' + options.punctuation;
  };

  // Static method default options.
  $.ScaleAndExecution.options = {
    punctuation: '.'
  };

  // Custom selector.
  $.expr[':'].ScaleAndExecution = function(elem) {
    // Is this element awesome?
    return $(elem).text().indexOf('awesome') !== -1;
  };

}(jQuery));
