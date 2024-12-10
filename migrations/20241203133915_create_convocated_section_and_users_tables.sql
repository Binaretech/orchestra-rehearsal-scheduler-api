-- +goose Up
-- +goose StatementBegin
create table public.concert_sections (
    id bigserial primary key,
    concert_id bigint not null constraint fk_concert_sections_concerts references public.concerts on delete cascade,
    section_id bigint not null constraint fk_concert_sections_sections references public.sections on delete cascade,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);

create table public.stands (
    id bigserial primary key,
    concert_section_id bigint not null constraint fk_stands_concert_sections references public.concert_sections on delete cascade,
    stand_number integer not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    constraint unique_section_stand unique (concert_section_id, stand_number)
);

create table public.stand_users (
    user_id bigint not null constraint fk_stand_users_users references public.users on delete cascade,
    stand_id bigint not null constraint fk_stand_users_stands references public.stands on delete cascade,
    primary key (stand_id, user_id)
);

create trigger update_concert_sections_timestamp before
update on public.concert_sections for each row execute procedure public.update_timestamp ();

create trigger update_stands_timestamp before
update on public.stands for each row execute procedure public.update_timestamp ();

create trigger update_stand_users_timestamp before
update on public.stand_users for each row execute procedure public.update_timestamp ();

create trigger update_families_timestamp before
update on public.families for each row execute procedure public.update_timestamp ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop trigger if exists update_stand_users_timestamp on public.stand_users;

drop table if exists public.stand_users;

drop trigger if exists update_stands_timestamp on public.stands;

drop table if exists public.stands;

drop trigger if exists update_concert_sections_timestamp on public.concert_sections;

drop table if exists public.concert_sections;

drop trigger if exists update_families_timestamp on public.families;

-- +goose StatementEnd
