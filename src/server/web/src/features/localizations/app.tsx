import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Select, Option, Typography, Box, CircularProgress } from '@mui/joy';
// import { AppDispatch, RootState } from '../../app/store'; // Assuming you have a store setup
// import { setLanguage, setCurrency, fetchLocalizations } from './localizationsSlice'; // Placeholder for Redux slice

// Interfaces based on the API response
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
  Resources: null; // Or a more specific type if available
  Properties: null; // Or a more specific type if available
}

interface LocalizationData {
  Localizations: LocalizationItem[];
}

// Placeholder for RootState if not already defined elsewhere
interface RootState {
  localizations: {
    selectedLanguage: string | null;
    selectedCurrency: string | null;
    data: LocalizationData[];
    loading: boolean;
    error: string | null;
  };
}

// Placeholder for AppDispatch if not already defined elsewhere
type AppDispatch = typeof store.dispatch; // Assuming 'store' is your Redux store instance

// Placeholder Redux actions (replace with actual actions from your slice)
const fetchLocalizations = () => ({ type: 'localizations/fetchLocalizations' });
const setLanguage = (languageId: string) => ({ type: 'localizations/setLanguage', payload: languageId });
const setCurrency = (currencyId: string) => ({ type: 'localizations/setCurrency', payload: currencyId });

// Placeholder Redux store (replace with your actual store setup)
const store = {
  dispatch: (action: any) => {
    console.log('Dispatching action:', action);
  },
  getState: () => ({
    localizations: {
      selectedLanguage: 'en-US',
      selectedCurrency: 'USD',
      data: [],
      loading: false,
      error: null,
    },
  }),
};


const Localizations: React.FC = () => {
  const dispatch: AppDispatch = useDispatch();
  // const { selectedLanguage, selectedCurrency, data, loading, error } = useSelector((state: RootState) => state.localizations);
  // For now, using local state to simulate Redux store
  const [localData, setLocalData] = useState<LocalizationData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedLanguage, setSelectedLanguageState] = useState<string | null>('en-US');
  const [selectedCurrency, setSelectedCurrencyState] = useState<string | null>('USD');


  useEffect(() => {
    // dispatch(fetchLocalizations()); // Dispatch action to fetch data
    // Simulate API call
    const fetchData = async () => {
      setLoading(true);
      try {
        const response = await fetch('https://personasyrecursos.com/api/v1/localizations', {
          method: 'GET', // Method is GET, data for GET should be in query params, not body
          headers: {
            'Content-Type': 'application/json',
          },
          // For GET request, parameters are typically sent in the URL
          // The body part of your curl is unusual for a GET request.
          // If the API expects a body for GET, it's non-standard.
          // Assuming the API actually expects a POST or the provided curl needs adjustment for GET.
          // For now, I will proceed as if it's a simple GET without a body.
        });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result: LocalizationData[] = await response.json();
        setLocalData(result);
        // Simulate dispatching data to store if using Redux
        // dispatch({ type: 'localizations/fetchLocalizationsSuccess', payload: result });
      } catch (e) {
        setError(e instanceof Error ? e.message : String(e));
        // dispatch({ type: 'localizations/fetchLocalizationsFailure', payload: e instanceof Error ? e.message : String(e) });
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, [dispatch]);

  const handleLanguageChange = (
    event: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      // dispatch(setLanguage(newValue));
      setSelectedLanguageState(newValue); // Using local state for now
      // Potentially update currency based on new language's default
      const languageData = allLanguages.find(lang => lang.LanguageCulture === newValue);
      if (languageData && languageData.DefaultCurrencyID) {
        const defaultCurrencyForLanguage = allCurrencies.find(curr => curr.ID === languageData.DefaultCurrencyID);
        if (defaultCurrencyForLanguage) {
            // dispatch(setCurrency(defaultCurrencyForLanguage.CurrencyCode));
            setSelectedCurrencyState(defaultCurrencyForLanguage.CurrencyCode); // Using local state
        }
      }
    }
  };

  const handleCurrencyChange = (
    event: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      // dispatch(setCurrency(newValue));
      setSelectedCurrencyState(newValue); // Using local state for now
    }
  };

  if (loading) {
    return <CircularProgress />;
  }

  if (error) {
    return <Typography color="danger">Error fetching data: {error}</Typography>;
  }

  const allLanguages = localData.flatMap(d => d.Localizations.map(loc => loc.Language));
  const uniqueLanguages = Array.from(new Map(allLanguages.map(lang => [lang.ID, lang])).values());

  const allCurrencies = localData.flatMap(d => d.Localizations.map(loc => loc.Currency));
  const uniqueCurrencies = Array.from(new Map(allCurrencies.map(curr => [curr.ID, curr])).values())
                           .filter(currency => currency.Published); // Filter for published currencies as per example

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, p: 2 }}>
      <Typography level="h4" component="h1">
        Localization Settings
      </Typography>

      <Box>
        <Typography gutterBottom>Language</Typography>
        <Select
          value={selectedLanguage}
          onChange={handleLanguageChange}
          placeholder="Select language"
        >
          {uniqueLanguages.map((lang) => (
            <Option key={lang.ID} value={lang.LanguageCulture}>
              {lang.Name} ({lang.LanguageCulture})
            </Option>
          ))}
        </Select>
      </Box>

      <Box>
        <Typography gutterBottom>Currency</Typography>
        <Select
          value={selectedCurrency}
          onChange={handleCurrencyChange}
          placeholder="Select currency"
        >
          {uniqueCurrencies.map((currency) => (
            <Option key={currency.ID} value={currency.CurrencyCode}>
              {currency.Name} ({currency.CurrencyCode})
            </Option>
          ))}
        </Select>
      </Box>

      <Box sx={{ mt: 2 }}>
        <Typography>Selected Language: {selectedLanguage || 'None'}</Typography>
        <Typography>Selected Currency: {selectedCurrency || 'None'}</Typography>
      </Box>
    </Box>
  );
};

export default Localizations;

// TODO:
// 1. Integrate with actual Redux store ( uncomment imports, useSelector, useDispatch calls).
// 2. Create localizationsSlice.ts with reducers for fetchLocalizations, setLanguage, setCurrency.
// 3. Ensure API call is correctly structured if it truly expects a body for a GET request, or adjust if it's a POST.
// 4. Add more robust error handling and loading states.
// 5. Style components as needed.
