import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react"

interface User {
    id: string;
    username: string;
    // Add other user properties as needed
  }
  
  interface AuthState {
    isAuthenticated: boolean;
    user: User | null;
    loading: 'idle' | 'pending' | 'succeeded' | 'failed';
    error: string | null;
    accessToken: string | null;
    refreshToken: string | null;
    message?: string; // Added message property
  }
  
  const initialState: AuthState = {
    isAuthenticated: false,
    user: null,
    loading: 'idle',
    error: null,
    accessToken: null, // Or load from localStorage if implementing token persistence
    refreshToken: null, // Or load from localStorage if implementing token persistence
  };

  // Define types for the login request and response
  interface LoginRequest {
    email?: string;
    password?: string;
    language?: string;
  }
  
  interface LoginResponse {
    // Adjust based on the actual API response structure
    // Assuming it returns a token and user information similar to loginSuccess payload
    accessToken: string;
    refreshToken: string; 
    // Example: if API returns data nested under a 'data' property
    // data: { token: string; user: User; };
    // success?: boolean;
    message?: string;
  }

  export const SigInApiSlice = createApi({
    baseQuery: fetchBaseQuery({ baseUrl: import.meta.env.VITE_API_URL }), // Corrected baseUrl
    reducerPath: "signInApi", // Changed reducerPath
    // tagTypes are removed as getQuotes is removed and login is a mutation
    endpoints: build => ({
      login: build.mutation<LoginResponse, LoginRequest>({
        query: ({ email, password, language }) => {
          const params = new URLSearchParams();
          
          if (email) params.append('email', email);
          if (password) params.append('password', password);
          if (language) params.append('language', language);
          else params.append('language', import.meta.env.VITE_LOCIZE_PROJECTID); // Default to 'en' if not provided
        
          return {
            url: `login?${params.toString()}`,
            method: 'POST',
            // Body can be added here if needed, but the curl example uses query params for POST
          };
        },
        // Example: invalidatesTags can be used if login should clear some cached data
        // invalidatesTags: [{ type: 'User', id: 'LIST' }], 
      }),
    }),
  })
  
  // Export the auto-generated hook for the login mutation
  export const { useLoginMutation } = SigInApiSlice;

  const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
      loginStart(state) {
        state.loading = 'pending';
        state.error = null;
      },
      loginSuccess(state, action: PayloadAction<{ user: User; accessToken: string; refreshToken: string  }>) {
        state.loading = 'succeeded';
        state.isAuthenticated = true;
        state.user = action.payload.user;
        state.accessToken = action.payload.accessToken;
        state.refreshToken = action.payload.refreshToken; // Assuming the API returns a refresh token
        // if (typeof window !== 'undefined') {
        //   localStorage.setItem('token', action.payload.token);
        // }
      },
      loginFailure(state, action: PayloadAction<string>) {
        state.loading = 'failed';
        state.isAuthenticated = false;
        state.user = null;
        state.error = action.payload;
        state.accessToken = null;
        state.refreshToken = null; // Reset tokens on failure
        state.message = action.payload; // Assuming the API returns a message on failure
        // if (typeof window !== 'undefined') {
        //   localStorage.removeItem('token');
        // }
      },
      logout(state) {
        state.isAuthenticated = false;
        state.user = null;
        state.accessToken = null;
        state.refreshToken = null; // Reset tokens on logout
        state.loading = 'idle';
        state.error = null;
        // if (typeof window !== 'undefined') {
        //   localStorage.removeItem('token');
        // }
      },
      // Example async action (thunk) would be defined separately or using createAsyncThunk
      // For simplicity, keeping it synchronous here.
      // You would typically dispatch loginStart, then make an API call,
      // then dispatch loginSuccess or loginFailure based on the API response.
    },
  });
  
  export const { loginStart, loginSuccess, loginFailure, logout } = authSlice.actions;
  
  // Selectors (optional, but good practice)
  export const selectIsAuthenticated = (state: { auth: AuthState }) => state.auth.isAuthenticated;
  export const selectUser = (state: { auth: AuthState }) => state.auth.user;
  export const selectAuthLoading = (state: { auth: AuthState }) => state.auth.loading;
  export const selectAuthError = (state: { auth: AuthState }) => state.auth.error;
  export const selectAuthToken = (state: { auth: AuthState }) => state.auth.accessToken;
  export const selectRefreshToken = (state: { auth: AuthState }) => state.auth.refreshToken
  export const selectMessage = (state: { auth: AuthState }) => state.auth.message
  
  export default authSlice;


