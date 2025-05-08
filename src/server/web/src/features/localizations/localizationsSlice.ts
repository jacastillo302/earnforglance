import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

// Interfaces based on the API response (can be shared or redefined here)
interface Language {
  ID: string;
  Name: string;
  LanguageCulture: string;
  UniqueSeoCode: string;
  FlagImageFileName: string;
  Rtl: boolean;
  LimitedToStores: boolean;
  DefaultCurrencyID: string;
  Published: boolean;
  DisplayOrder: number;
}

interface Currency {
  ID: string;
  Name: string;
  CurrencyCode: string;
  Rate: number;
  DisplayLocale: string;
  CustomFormatting: string;
  LimitedToStores: boolean;
  Published: boolean;
  DisplayOrder: number;
  CreatedOnUtc: string;
  UpdatedOnUtc: string;
  RoundingTypeID: number;
}

interface LocalizationItem {
  Language: Language;
  Currency: Currency;
  Resources: null; // Or a more specific type
  Properties: null; // Or a more specific type
}

export interface LocalizationData {
  Localizations: LocalizationItem[];
}

// Define the API slice
export const localizationsApi = createApi({
  reducerPath: 'localizationsApi',
  baseQuery: fetchBaseQuery({ baseUrl: 'https://personasyrecursos.com/api/v1/' }),
  endpoints: (builder) => ({
    getLocalizations: builder.query<LocalizationData[], void>({
      query: () => ({
        url: 'localizations',
        method: 'POST',
        body: JSON.stringify({
          ID: "",
          Filters: [], // Modified to be an array directly
          Sort: "asc",
          Limit: 10,
          Page: 1,
          Lang: "",
          AllowComments: false,
          Content: ["currency"], // Modified to be an array directly
        }),
      }),
    }),
  }),
});

// Export hooks for usage in functional components
export const { useGetLocalizationsQuery } = localizationsApi;

interface LocalizationsState {
  selectedLanguage: string | null;
  selectedCurrency: string | null;
  data: LocalizationData[];
}

const initialState: LocalizationsState = {
  selectedLanguage: 'en-US', // Default language
  selectedCurrency: 'USD',   // Default currency
  data: [],
};

const localizationsSlice = createSlice({
  name: 'localizations',
  initialState,
  reducers: {
    setLanguage: (state, action: PayloadAction<string>) => {
      state.selectedLanguage = action.payload;
      // Optionally, update currency based on the new language's default
      const languageData = state.data
        .flatMap(d => d.Localizations.map(loc => loc.Language))
        .find(lang => lang.LanguageCulture === action.payload);

      if (languageData && languageData.DefaultCurrencyID) {
        const defaultCurrencyForLanguage = state.data
          .flatMap(d => d.Localizations.map(loc => loc.Currency))
          .find(curr => curr.ID === languageData.DefaultCurrencyID);
        if (defaultCurrencyForLanguage) {
          state.selectedCurrency = defaultCurrencyForLanguage.CurrencyCode;
        }
      }
    },
    setCurrency: (state, action: PayloadAction<string>) => {
      state.selectedCurrency = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addMatcher(
      localizationsApi.endpoints.getLocalizations.matchFulfilled,
      (state, action: PayloadAction<LocalizationData[]>) => {
        state.data = action.payload;
      }
    );
  },
});

export const { setLanguage, setCurrency } = localizationsSlice.actions;
export default localizationsSlice;
