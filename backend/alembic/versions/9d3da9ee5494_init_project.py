"""init project

Revision ID: 9d3da9ee5494
Revises: 
Create Date: 2024-05-09 11:58:20.179636

"""
from typing import Sequence, Union
from sqlalchemy.engine.reflection import Inspector

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = '9d3da9ee5494'
down_revision: Union[str, None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    conn = op.get_bind()
    inspector = Inspector.from_engine(conn)
    tables = inspector.get_table_names()

    if 'articles' not in tables:
        op.create_table(
            'articles',
            sa.Column('id', sa.Integer(), nullable=False),
            sa.Column('title', sa.String(), nullable=False),
            sa.Column('content', sa.String(), nullable=False),
            sa.Column('published', sa.Boolean(), server_default='TRUE', nullable=False),
            sa.Column('created_at', sa.TIMESTAMP(timezone=True), server_default=sa.text('now()'), nullable=False),
            sa.Column('updated_at', sa.TIMESTAMP(timezone=True), server_default=sa.text('now()'), nullable=False),
            sa.PrimaryKeyConstraint('id')
        )


def downgrade() -> None:
    op.drop_table('articles')
