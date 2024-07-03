import React from 'react';
import { Route, Redirect } from 'react-router-dom';
// for routes
const ProtectedRoute = ({ component: Component, roles, ...rest }) => {
  const userRoles = JSON.parse(sessionStorage.getItem('userRoles') || '[]')

  return (
    <Route
      {...rest}
      render={props => {
        console.log(userRoles)
        if (!userRoles) {
          // If no roles found, redirect to login
          return <Redirect to="/" />;
        }

        // Check if user has the required role
        const hasRequiredRole = roles.some(role => userRoles.includes(role));
        console.log("Has resuired role: ",hasRequiredRole)
        
        if (!hasRequiredRole) {
          // If user doesn't have the required role, redirect to error page
          return <Redirect to="/*" />;
        }

        // If user has the required role, render the component
        return <Component {...props} />;
      }}
    />
  );
};

export default ProtectedRoute;