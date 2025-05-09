import React, { useState, useEffect } from 'react';
import { Select, Option, Typography, Box, CircularProgress } from '@mui/joy';
import { useGetLocalizationsQuery } from './localizationsSlice'; // Import the hook and interface

const Localizations: React.FC = () => {
  // Attempt to load from localStorage initially
  const [cachedLanguages, setCachedLanguages] = useState<any[] | null>(() => {
    try {
      const item = localStorage.getItem('appCachedLanguages');
      if (item) {
        return JSON.parse(item);
      }
      return null;
    } catch (error) {
      console.error("Error reading appCachedLanguages from localStorage", error);
      // Attempt to remove corrupted item
      try {
        localStorage.removeItem('appCachedLanguages');
      } catch (removeError) {
        console.error("Error removing corrupted appCachedLanguages from localStorage", removeError);
      }
      return null;
    }
  });
  const [cachedCurrencies, setCachedCurrencies] = useState<any[] | null>(() => {
    try {
      const item = localStorage.getItem('appCachedCurrencies');
      if (item) {
        return JSON.parse(item);
      }
      return null;
    } catch (error) {
      console.error("Error reading appCachedCurrencies from localStorage", error);
      // Attempt to remove corrupted item
      try {
        localStorage.removeItem('appCachedCurrencies');
      } catch (removeError) {
        console.error("Error removing corrupted appCachedCurrencies from localStorage", removeError);
      }
      return null;
    }
  });

  const [selectedLanguage, setSelectedLanguageState] = useState<string | null>(null);
  const [selectedCurrency, setSelectedCurrencyState] = useState<string | null>('USD'); // Default USD

  // API call, skipped if data is already cached from localStorage
  const { data: localData, error, isLoading } = useGetLocalizationsQuery(
    undefined,
    { skip: !!(cachedLanguages && cachedCurrencies && cachedLanguages.length > 0 && cachedCurrencies.length > 0) } // Skip if both have data
  );

  // Effect to process data from API and save to localStorage
  useEffect(() => {
    if (localData) {
      // Process languages from API data
      const allLanguagesApi = localData.flatMap(entry =>
        entry.Localizations ? entry.Localizations.map(loc => loc.Language) : []
      );
      const processedLanguagesApi = allLanguagesApi
        .filter(lang =>
          lang &&
          lang.ID != null &&
          lang.LanguageCulture != null &&
          lang.Name != null &&
          lang.FlagImageFileName != null &&
          lang.UniqueSeoCode != null
        )
        .map(lang => ({
          ...lang,
          UniqueSeoCode: lang.UniqueSeoCode.toUpperCase()
        }));
      const uniqueLanguagesApi = Array.from(
        new Map(processedLanguagesApi.map(lang => [lang.ID, lang])).values()
      );
      setCachedLanguages(uniqueLanguagesApi);
      try {
        localStorage.setItem('appCachedLanguages', JSON.stringify(uniqueLanguagesApi));
      } catch (e) {
        console.error("Error saving languages to localStorage", e);
      }

      // Process currencies from API data
      const allCurrenciesApi = localData.flatMap(entry =>
        entry.Localizations ? entry.Localizations.map(loc => loc.Currency) : []
      );
      const uniqueCurrenciesApi = Array.from(
        new Map(
          allCurrenciesApi
            .filter(curr => curr && curr.ID != null && curr.CurrencyCode != null && curr.Name != null)
            .map(curr => [curr.ID, curr])
        ).values()
      );
      setCachedCurrencies(uniqueCurrenciesApi);
      try {
        localStorage.setItem('appCachedCurrencies', JSON.stringify(uniqueCurrenciesApi));
      } catch (e) {
        console.error("Error saving currencies to localStorage", e);
      }
    }
  }, [localData]);

  // Effect to set default selected language when cachedLanguages are available/updated
  useEffect(() => {
    if (cachedLanguages && cachedLanguages.length > 0) {
      // If selectedLanguage is not set OR is not in the current list of cachedLanguages
      if (!selectedLanguage || !cachedLanguages.find(lang => lang.UniqueSeoCode === selectedLanguage)) {
        const defaultLang = cachedLanguages.find(lang => lang.LanguageCulture === 'en-US');
        if (defaultLang) {
          setSelectedLanguageState(defaultLang.UniqueSeoCode);
        } else {
          setSelectedLanguageState(cachedLanguages[0].UniqueSeoCode); // Fallback to the first language
        }
      }
    } else if (cachedLanguages && cachedLanguages.length === 0) {
      setSelectedLanguageState(null); // No languages available
    } else if (!cachedLanguages) { // If cachedLanguages itself is null
      setSelectedLanguageState(null);
    }
  }, [cachedLanguages, selectedLanguage]);


  const handleLanguageChange = (
    _: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      setSelectedLanguageState(newValue);
    }
  };

  const handleCurrencyChange = (
    _: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      setSelectedCurrencyState(newValue);
    }
  };

  if (isLoading) { // API is loading
    return <CircularProgress />;
  }

  if (error) {
    return <Typography color="danger">Error fetching data: {('data' in error ? JSON.stringify(error.data) : "error.message") || 'Unknown error'}</Typography>;
  }

  // If not loading from API and no cached languages after attempting to load/fetch
  if (!cachedLanguages || cachedLanguages.length === 0) {
    return <Typography>No language data available.</Typography>;
  }

  return (
    <Box sx={{ display: 'flex', flexDirection: 'row', gap: 2, p: 2, alignItems: 'center' }}>
      <Box sx={{ flexGrow: 1 }}>
        <Select
          value={selectedLanguage}
          onChange={handleLanguageChange}
          placeholder="Select language"
          disabled={!cachedLanguages || cachedLanguages.length === 0}
        >
          {cachedLanguages && cachedLanguages.map((lang) => (
            <Option key={lang.ID} value={lang.UniqueSeoCode}>
              <img
                src={`${import.meta.env.VITE_MEDIA_URL}images/${lang.FlagImageFileName}`}
                alt={`${lang.Name} flag`}
                style={{ width: '20px', height: 'auto', marginRight: '8px', verticalAlign: 'middle' }}
              />
              {lang.UniqueSeoCode}
            </Option>
          ))}
        </Select>
      </Box>

      <Box sx={{ flexGrow: 1 }}>
        <Select
          value={selectedCurrency}
          onChange={handleCurrencyChange}
          placeholder="Select currency"
          disabled={!cachedCurrencies || cachedCurrencies.length === 0}
        >
          {cachedCurrencies && cachedCurrencies.map((currency) => ( // Added null check for cachedCurrencies
            <Option key={currency.ID} value={currency.CurrencyCode}>
              {currency.CurrencyCode}
            </Option>
          ))}
        </Select>
      </Box>
    </Box>
  );
};

export default Localizations;
