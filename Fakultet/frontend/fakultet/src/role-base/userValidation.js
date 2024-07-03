import { useState, useEffect } from 'react';
// for use in code
const useRoles = () => {
  const [userRoles, setUserRoles] = useState([]);

  useEffect(() => {
    const roles = JSON.parse(sessionStorage.getItem('userRoles') || '[]');
    setUserRoles(roles);
  }, []);

  const hasRole = (role) => userRoles.includes(role);

  return { userRoles, hasRole };
};

export default useRoles;
