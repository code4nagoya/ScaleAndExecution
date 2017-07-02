/*jshint node:true */
module.exports = function(grunt) {
	grunt.loadNpmTasks("grunt-contrib-jshint");
	grunt.loadNpmTasks("grunt-contrib-qunit");
	grunt.loadNpmTasks("grunt-contrib-uglify");
	grunt.loadNpmTasks("grunt-contrib-cssmin");
	grunt.loadNpmTasks("grunt-qunit-junit");

	grunt.initConfig({
		qunit : {
			all : [ "tests/group.html" ]
		},
		qunit_junit : {
			options : {
				dest : 'reports/'
			}
		},

		uglify : {
			options : {},
			externalLibraries : {

			},
			app : {
				files : {

				}
			}
		},

		cssmin : {

			combine:{
				files:{
					'build/app.min.css' : ['css/**.css']
				}
			}
		},

		jshint : {
			options : {
				reportOutpu : 'reports/jshint.txt'
			}
		}
	});

	grunt.registerTask("test", [ "qunit" ]);
	grunt.registerTask("default", [ "jshint", "uglify", "cssmin",
			"qunit_junit", "test" ]);

};
