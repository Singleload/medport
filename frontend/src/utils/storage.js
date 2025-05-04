/**
 * Storage utility for localStorage and sessionStorage operations
 */
const storage = {
    /**
     * Sets an item in localStorage with optional expiration time
     * @param {string} key - Storage key
     * @param {*} value - Value to store
     * @param {number} expiryInMinutes - Optional expiration time in minutes
     */
    setLocalItem(key, value, expiryInMinutes = null) {
      try {
        const item = {
          value: value,
          timestamp: new Date().getTime()
        };
        
        // Add expiration if specified
        if (expiryInMinutes) {
          item.expiry = expiryInMinutes * 60 * 1000; // Convert to milliseconds
        }
        
        localStorage.setItem(key, JSON.stringify(item));
      } catch (error) {
        console.error('Error setting localStorage item:', error);
      }
    },
    
    /**
     * Gets an item from localStorage with expiration check
     * @param {string} key - Storage key
     * @returns {*} - Stored value or null if expired/not found
     */
    getLocalItem(key) {
      try {
        const itemStr = localStorage.getItem(key);
        
        // Return null if item doesn't exist
        if (!itemStr) {
          return null;
        }
        
        const item = JSON.parse(itemStr);
        const now = new Date().getTime();
        
        // Check if the item has expired
        if (item.expiry && now > item.timestamp + item.expiry) {
          // Item has expired, remove it
          localStorage.removeItem(key);
          return null;
        }
        
        return item.value;
      } catch (error) {
        console.error('Error getting localStorage item:', error);
        return null;
      }
    },
    
    /**
     * Removes an item from localStorage
     * @param {string} key - Storage key
     */
    removeLocalItem(key) {
      try {
        localStorage.removeItem(key);
      } catch (error) {
        console.error('Error removing localStorage item:', error);
      }
    },
    
    /**
     * Sets an item in sessionStorage
     * @param {string} key - Storage key
     * @param {*} value - Value to store
     */
    setSessionItem(key, value) {
      try {
        sessionStorage.setItem(key, JSON.stringify(value));
      } catch (error) {
        console.error('Error setting sessionStorage item:', error);
      }
    },
    
    /**
     * Gets an item from sessionStorage
     * @param {string} key - Storage key
     * @returns {*} - Stored value or null if not found
     */
    getSessionItem(key) {
      try {
        const item = sessionStorage.getItem(key);
        return item ? JSON.parse(item) : null;
      } catch (error) {
        console.error('Error getting sessionStorage item:', error);
        return null;
      }
    },
    
    /**
     * Removes an item from sessionStorage
     * @param {string} key - Storage key
     */
    removeSessionItem(key) {
      try {
        sessionStorage.removeItem(key);
      } catch (error) {
        console.error('Error removing sessionStorage item:', error);
      }
    },
    
    /**
     * Clears all items from localStorage
     */
    clearLocalStorage() {
      try {
        localStorage.clear();
      } catch (error) {
        console.error('Error clearing localStorage:', error);
      }
    },
    
    /**
     * Clears all items from sessionStorage
     */
    clearSessionStorage() {
      try {
        sessionStorage.clear();
      } catch (error) {
        console.error('Error clearing sessionStorage:', error);
      }
    }
  };
  
  export default storage;