import React, { useState, useEffect } from 'react';
import { Select, Option, Typography, Box, CircularProgress } from '@mui/joy';
import { useGetLocalizationsQuery } from './localizationsSlice'; // Import the hook and interface

const Localizations: React.FC = () => {
  const { data: localData, error, isLoading: loading } = useGetLocalizationsQuery(); // Use the hook
  const [selectedLanguage, setSelectedLanguageState] = useState<string | null>(null); // Changed initial state
  const [selectedCurrency, setSelectedCurrencyState] = useState<string | null>('USD');
  const [cachedLanguages, setCachedLanguages] = useState<any[] | null>(null);
  const [cachedCurrencies, setCachedCurrencies] = useState<any[] | null>(null);

  useEffect(() => {
    if (localData) {
      if (!cachedLanguages) {
        const allLanguages = localData.flatMap(entry =>
          entry.Localizations ? entry.Localizations.map(loc => loc.Language) : []
        );
        const processedLanguages = allLanguages
          .filter(lang =>
            lang &&
            lang.ID != null &&
            lang.LanguageCulture != null &&
            lang.Name != null &&
            lang.FlagImageFileName != null && // Added check
            lang.UniqueSeoCode != null        // Added check
          )
          .map(lang => ({
            ...lang,
            UniqueSeoCode: lang.UniqueSeoCode.toUpperCase() // Uppercase UniqueSeoCode
          }));

        const uniqueLanguages = Array.from(
          new Map(processedLanguages.map(lang => [lang.ID, lang])).values()
        );
        setCachedLanguages(uniqueLanguages);

        // Set initial selected language
        if (uniqueLanguages.length > 0) {
          const defaultLang = uniqueLanguages.find(lang => lang.LanguageCulture === 'en-US');
          if (defaultLang) {
            setSelectedLanguageState(defaultLang.UniqueSeoCode); // Already uppercased
          } else {
            // Fallback to the first language in the list if 'en-US' not found
            setSelectedLanguageState(uniqueLanguages[0].UniqueSeoCode); // Already uppercased
          }
        } else {
          setSelectedLanguageState(null);
        }
      }

      if (!cachedCurrencies) {
        const allCurrencies = localData.flatMap(entry =>
          entry.Localizations ? entry.Localizations.map(loc => loc.Currency) : []
        );
        const uniqueCurrencies = Array.from(
          new Map(
            allCurrencies
              .filter(curr => curr && curr.ID != null && curr.CurrencyCode != null && curr.Name != null)
              .map(curr => [curr.ID, curr])
          ).values()
        );
        setCachedCurrencies(uniqueCurrencies);
      }
    }
  }, [localData, cachedLanguages, cachedCurrencies]);

  const handleLanguageChange = (
    _: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      setSelectedLanguageState(newValue); // Using local state for now
      // Removed logic to auto-select currency based on language
    }
  };

  const handleCurrencyChange = (
    _: React.SyntheticEvent | null,
    newValue: string | null,
  ) => {
    if (newValue) {
      setSelectedCurrencyState(newValue); // Using local state for now
    }
  };

  if (loading) {
    return <CircularProgress />;
  }

  if (error) {
    return <Typography color="danger">Error fetching data: {('data' in error ? JSON.stringify(error.data) : "error.message") || 'Unknown error'}</Typography>;
  }

  if (!localData || !cachedLanguages || !cachedCurrencies) {
    if (!localData || !cachedCurrencies) {
      return <Typography>No data available.</Typography>;
    }
  }

  return (
    <Box sx={{ display: 'flex', flexDirection: 'row', gap: 2, p: 2, alignItems: 'center' }}>
      <Box sx={{ flexGrow: 1 }}>
        <Select
          value={selectedLanguage}
          onChange={handleLanguageChange}
          placeholder="Select language"
          disabled={!cachedLanguages || cachedLanguages.length === 0} // Disable if no languages
        >
          {cachedLanguages && cachedLanguages.map((lang) => (
            <Option key={lang.ID} value={lang.UniqueSeoCode}> {/* lang.UniqueSeoCode is already uppercased */}
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
        >
          {cachedCurrencies.map((currency) => (
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
