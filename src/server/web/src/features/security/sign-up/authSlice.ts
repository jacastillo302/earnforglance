import { createSlice, PayloadAction } from '@reduxjs/toolkit';


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
    token: string | null;
  }
  
  const initialState: AuthState = {
    isAuthenticated: false,
    user: null,
    loading: 'idle',
    error: null,
    token: null, // Or load from localStorage if implementing token persistence
  };
  
  const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
      loginStart(state) {
        state.loading = 'pending';
        state.error = null;
      },
      loginSuccess(state, action: PayloadAction<{ user: User; token: string }>) {
        state.loading = 'succeeded';
        state.isAuthenticated = true;
        state.user = action.payload.user;
        state.token = action.payload.token;
        // if (typeof window !== 'undefined') {
        //   localStorage.setItem('token', action.payload.token);
        // }
      },
      loginFailure(state, action: PayloadAction<string>) {
        state.loading = 'failed';
        state.isAuthenticated = false;
        state.user = null;
        state.error = action.payload;
        state.token = null;
        // if (typeof window !== 'undefined') {
        //   localStorage.removeItem('token');
        // }
      },
      logout(state) {
        state.isAuthenticated = false;
        state.user = null;
        state.token = null;
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
  export const selectAuthToken = (state: { auth: AuthState }) => state.auth.token;
  
  export default authSlice;
  
  // Example of how you might define an async thunk for login (requires redux-thunk middleware)
  /*
  import { AppThunk } from '../../app/store'; // Assuming your store setup exports AppThunk type
  
  export const loginUser = (
    credentials: { username: string; password: string }
  ): AppThunk => async (dispatch) => {
    try {
      dispatch(loginStart());
      // const response = await api.post('/login', credentials); // Replace with your API call
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1000));
      const simulatedResponse = {
        user: { id: '1', username: credentials.username },
        token: 'fake-jwt-token',
      };
      dispatch(loginSuccess(simulatedResponse));
    } catch (error: any) {
      dispatch(loginFailure(error.message || 'Failed to login'));
    }
  };
  */
  
  
