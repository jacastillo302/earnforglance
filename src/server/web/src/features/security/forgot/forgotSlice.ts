import { createSlice, PayloadAction } from '@reduxjs/toolkit';

// Define a type for the user data
interface User {
  id: string;
  name: string;
  email: string;
  // Add other user properties as needed
}

// Define the initial state for the auth slice
interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  token: string | null;
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | null;
}

const initialState: AuthState = {
  isAuthenticated: false,
  user: null,
  token: null,
  status: 'idle',
  error: null,
};

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    loginStart(state) {
      state.status = 'loading';
    },
    loginSuccess(state, action: PayloadAction<{ user: User; token: string }>) {
      state.isAuthenticated = true;
      state.user = action.payload.user;
      state.token = action.payload.token;
      state.status = 'succeeded';
      state.error = null;
    },
    loginFailure(state, action: PayloadAction<string>) {
      state.isAuthenticated = false;
      state.user = null;
      state.token = null;
      state.status = 'failed';
      state.error = action.payload;
    },
    logout(state) {
      state.isAuthenticated = false;
      state.user = null;
      state.token = null;
      state.status = 'idle';
      state.error = null;
    },
    // Example action for updating user profile
    updateUserProfile(state, action: PayloadAction<Partial<User>>) {
      if (state.user) {
        state.user = { ...state.user, ...action.payload };
      }
    },
    // Example action for password recovery request
    passwordRecoveryStart(state) {
      state.status = 'loading';
    },
    passwordRecoverySuccess(state) {
      state.status = 'succeeded'; // Or a specific status like 'recoveryEmailSent'
      state.error = null;
    },
    passwordRecoveryFailure(state, action: PayloadAction<string>) {
      state.status = 'failed';
      state.error = action.payload;
    },
  },
});

export const {
  loginStart,
  loginSuccess,
  loginFailure,
  logout,
  updateUserProfile,
  passwordRecoveryStart,
  passwordRecoverySuccess,
  passwordRecoveryFailure,
} = authSlice.actions;

// Selectors (optional, but good practice)
// Replace RootState with your actual root state type if different
interface RootState {
  auth: AuthState;
}

export const selectIsAuthenticated = (state: RootState) => state.auth.isAuthenticated;
export const selectUser = (state: RootState) => state.auth.user;
export const selectAuthStatus = (state: RootState) => state.auth.status;
export const selectAuthError = (state: RootState) => state.auth.error;

export default authSlice.reducer;
