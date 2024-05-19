SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: cardnetwork_enum; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.cardnetwork_enum AS ENUM (
    'MasterCard',
    'Amex',
    'Visa'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: merchants; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.merchants (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transactions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_account_id uuid NOT NULL,
    card_network public.cardnetwork_enum NOT NULL,
    merchant_id text NOT NULL,
    merchant_name text NOT NULL,
    total numeric(10,2) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


--
-- Name: user_accounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_accounts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    card_network public.cardnetwork_enum NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    email text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


--
-- Name: merchants merchants_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.merchants
    ADD CONSTRAINT merchants_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: user_accounts user_accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_accounts
    ADD CONSTRAINT user_accounts_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_merchants__name; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_merchants__name ON public.merchants USING btree (name);


--
-- Name: idx_transactions__created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_transactions__created_at ON public.transactions USING btree (created_at);


--
-- Name: idx_transactions__merchant_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_transactions__merchant_id ON public.transactions USING btree (merchant_id);


--
-- Name: idx_transactions__user_account_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_transactions__user_account_id ON public.transactions USING btree (user_account_id);


--
-- Name: idx_transactions_card_network; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_transactions_card_network ON public.transactions USING btree (card_network);


--
-- Name: idx_user_accounts__card_network; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_accounts__card_network ON public.user_accounts USING btree (card_network);


--
-- Name: idx_user_accounts__user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_accounts__user_id ON public.user_accounts USING btree (user_id);


--
-- Name: idx_users__email; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_users__email ON public.users USING btree (email);


--
-- Name: transactions transactions_user_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_user_account_id_fkey FOREIGN KEY (user_account_id) REFERENCES public.user_accounts(id);


--
-- Name: user_accounts user_accounts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_accounts
    ADD CONSTRAINT user_accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20240518232133'),
    ('20240518232142'),
    ('20240518232143'),
    ('20240518232149'),
    ('20240519162851');
