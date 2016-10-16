var express    = require('express'),
    requireDir = require('require-dir'),
    assert     = require('assert');

var config      = require('./config.js'),
    controllers = requireDir('./controllers');

module.exports = function(db, passport) {
    router = express.Router();

    // Helper for each call
    var checkAuth = function (req, res, next) {
        if (req.isAuthenticated()) {
            return next();
        } else {
            req.flash('error', 'Unauthenticated request');
            res.status(401).redirect('/loginPage');
        };
    };

    // Wrapper for adding all function calls
    var addRoute = function(method, route, handler) {
        router[method]('/' + route,
               checkAuth,
               handler);
    };

    // Custom passport authentication handler
    var isLocalAuthenticated = function(req, res, next) {
        passport.authenticate('local', function(err, user, info) {
            if (err) { return next(err); }

            // user will be set to false, if not authenticated
            if (!user) {
                res.status(401).json(info); //info contains the error message
            } else {
                // if user authenticated maintain the session
                req.logIn(user, function() {
                    res.redirect(config.api + '/test');
                })
            };
        })(req, res, next);
    }

    // Log In, Log Out status
    router.post('/login', isLocalAuthenticated);
    router.get('/logout', function(req, res) {
        req.logout();
        res.redirect('/');
    });


    // ==============
    // All the routes
    // ==============
    addRoute('get', '', function(req, res) {
        res.send('Hello from the other side');
    });

    addRoute('get',  'findUser', controllers.user.findUser(db));

    // TODO: Only for testing passport.js
    // Needs to be removed in production
    // #security, #important
    router.get('/findUser2', controllers.user.findUser(db));

    addRoute('post', 'update/first', controllers.user.firstLogin(db));
    addRoute('post', 'update/pass', controllers.user.changePassword(db));
    addRoute('post', 'info/login', controllers.user.getInfoOnLogin(db));

    return router;
};